package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/crewjam/saml"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// Once up and running, can register with https://samltest.id/upload.php
// Run `make metadata` to gather the saml meta required for the saml test
func main() {
	//getAWSServiceProvider()
	samlSP, err := startServiceProvider()
	if err != nil {
		log.Fatal(err)
	}

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

	e := xml.NewEncoder(os.Stdout)
	e.Indent(" ", "  ")
	authReq.MarshalXML(e, xml.StartElement{})

	// relayState, err := samlSP.RequestTracker.TrackRequest(http.ResponseWriter, http.Request, authReq.ID)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Set up service provider as a server
	http.Handle("/hello", samlSP.RequireAccount(http.HandlerFunc(hello)))
	http.Handle("/saml/", samlSP)
	http.ListenAndServe(":8000", nil)
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
