package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type ParseJSON struct {
	Region           string
	AccountId        string
	Environment      string
	SourceBucketName string
	S3KmsKeyName     string
}

func main() {
	file := flag.String("file", "None", "What file to parse")
	region := flag.String("region", "us-east-1", "What region to include in the JSON file")
	accountId := flag.String("accountId", "0", "The account number to include in JSON file")
	environment := flag.String("environment", "None", "The environment to deploy to in JSON file")
	sourceBucket := flag.String("sourceBucket", "None", "The source bucket to insert into the JSON file")
	s3Kms := flag.String("s3Kms", "None", "The S3 KMS Key to insert into the S3 bucket")
	flag.Parse()

	if *file == "None" {
		fmt.Println("Usage: jsonencode -file=\"<file-name>\" -region=\"<aws-region>\" -accountId=\"<account-id>\" -environment=\"<environment>\" -sourceBucket=\"<source-bucket-name>\" -s3Kms=\"<s3-kms-key-name>\"")
	}

	jsonTest := new(ParseJSON)
	jsonTest.Region = *region
	jsonTest.AccountId = *accountId
	jsonTest.Environment = *environment
	jsonTest.SourceBucketName = *sourceBucket
	jsonTest.S3KmsKeyName = *s3Kms

	Encode(*file, jsonTest)

	return
}

func Encode(file string, jsonTest *ParseJSON) {
	jsonExample, err := ioutil.ReadFile(file)

	template, err := template.New("InputRequest").Parse(string(jsonExample))
	CheckError(err)

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	err = template.Execute(doc, jsonTest)
	CheckError(err)

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	CheckError(err)

	err = ioutil.WriteFile("output.json", []byte(doc.String()), 0644)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
