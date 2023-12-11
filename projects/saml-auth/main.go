package main

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/crewjam/saml"
	"github.com/crewjam/saml/samlsp"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
		return
	}

	samlResponse := r.Form.Get("SAMLResponse")
	if samlResponse == "" {
		fmt.Fprint(w, "No SAMLResponse found\n")
	}

	decodedResponse, err := url.QueryUnescape(samlResponse)
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
		return
	}

	fmt.Fprintf(w, "SAMLResponse: %s\n", decodedResponse)
	for _, cookie := range r.Cookies() {
		fmt.Fprintf(w, "Received cookie: %s = %s\n", cookie.Name, cookie.Value)
		decoded, err := base64.StdEncoding.DecodeString(cookie.Value)
		if err != nil {
			fmt.Fprintf(w, "Error: %v\n", err)
			return
		}
		fmt.Println("cookie token decoded:", decoded)
	}
	if len(r.Cookies()) == 0 {
		fmt.Fprint(w, "No cookies received\n")
	}
}

// Once up and running, can register with https://samltest.id/upload.php
// Run `make metadata` to gather the saml meta required for the saml test
func main() {
	//getAWSServiceProvider()
	samlSP, err := startServiceProvider()
	if err != nil {
		log.Fatal(err)
	}

	// Set up service provider as a server
	http.Handle("/hello", samlSP.RequireAccount(http.HandlerFunc(hello)))
	http.Handle("/saml/", samlSP)
	http.ListenAndServe(":8000", nil)
}

func startServiceProvider() (*samlsp.Middleware, error) {
	keyPair, err := tls.LoadX509KeyPair("myservice.cert", "myservice.key")
	if err != nil {
		return nil, err
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		return nil, err
	}

	idpMetadataURL, err := url.Parse("https://samltest.id/saml/idp")
	if err != nil {
		return nil, err
	}

	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient, *idpMetadataURL)
	if err != nil {
		return nil, err
	}

	rootURL, err := url.Parse("http://localhost:8000")
	if err != nil {
		return nil, err
	}

	return samlsp.New(samlsp.Options{
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})

}

// aws iam list-saml-providers
// aws iam get-saml-provider --saml-provider-arn
func getAWSServiceProvider() {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("cn-north-1"),
		config.WithSharedConfigProfile(""),
	)
	if err != nil {
		log.Fatal(err)
	}

	svc := iam.NewFromConfig(cfg)
	output, err := svc.GetSAMLProvider(context.Background(), &iam.GetSAMLProviderInput{
		SAMLProviderArn: aws.String(""),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*output.SAMLMetadataDocument)
}

func runRawRequest(samlSP *samlsp.Middleware) {
	var b, bl string
	if samlSP.Binding == "" {
		b = saml.HTTPRedirectBinding
		bl = samlSP.ServiceProvider.GetSSOBindingLocation(b)
		if bl == "" {
			b = saml.HTTPPostBinding
			bl = samlSP.ServiceProvider.GetSSOBindingLocation(b)
		}
	} else {
		b = samlSP.Binding
		bl = samlSP.ServiceProvider.GetSSOBindingLocation(b)
	}

	log.Println("ACS URL:", samlSP.ServiceProvider.AcsURL.Path)
	log.Println("binding:", b)
	log.Println("binding location:", bl)
	log.Println("result binding:", samlSP.ResponseBinding)

	authReq, err := samlSP.ServiceProvider.MakeAuthenticationRequest(bl, b, samlSP.ResponseBinding)
	if err != nil {
		log.Fatal(err)
	}

	// Print the XML to the console
	e := xml.NewEncoder(os.Stdout)
	e.Indent(" ", "  ")
	authReq.MarshalXML(e, xml.StartElement{})

	mxml, err := xml.Marshal(authReq)
	if err != nil {
		log.Fatal(err)
	}

	samlReq := base64.StdEncoding.EncodeToString(mxml)
	// Create an HTTP POST request with the SAMLRequest as a form value
	data := url.Values{
		"SAMLRequest": {samlReq},
	}

	httpClient := &http.Client{}
	r, err := http.NewRequest("POST", authReq.Destination, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	log.Println("sending request to", authReq.Destination)
	// Send the HTTP request using an HTTP client
	resp, err := httpClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println()

	log.Println("response status:", resp.Status)
	log.Println("response headers:", resp.Header)
	log.Println("response body:", resp.Body)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	formValues, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		log.Fatal(err)
	}

	samlResponse := formValues.Get("SAMLResponse")

	// Decode the SAMLResponse from base64
	decodedResponse, err := base64.StdEncoding.DecodeString(samlResponse)
	if err != nil {
		log.Fatal(err)
	}

	assertion, err := samlSP.ServiceProvider.ParseXMLResponse(decodedResponse, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print the attributes returned in the SAML Assertion
	for _, attributeStatement := range assertion.AttributeStatements {
		for _, attribute := range attributeStatement.Attributes {
			log.Printf("Attribute %s has value %s", attribute.Name, attribute.Values[0].Value)
		}
	}
}
