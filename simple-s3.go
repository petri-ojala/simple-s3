package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var awsRegion = flag.String("region", "eu-north-1", "AWS Region")
var s3Bucket = flag.String("bucket", "", "S3 bucket")
var s3Key = flag.String("name", "", "S3 object name")
var inputFile = flag.String("file", "", "Input filename")

var Usage = func() {
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if *awsRegion == "" || *s3Bucket == "" || *s3Key == "" || *inputFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	// AWS credentials
	creds := credentials.NewEnvCredentials()

	// Read input file
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Upload gzipped test result json to S3
	s3Config := &aws.Config{
		Credentials: creds,
		Region:      aws.String(*awsRegion),
		/*	DisableSSL:       aws.Bool(false,
			S3ForcePathStyle: aws.Bool(true),*/
	}

	s3Session := session.New(s3Config)

	_, err = s3.New(s3Session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(*s3Bucket),
		Key:    aws.String(*s3Key),
		Body:   bytes.NewReader(b),
	})
	if err != nil {
		log.Fatal(err)
	}
}
