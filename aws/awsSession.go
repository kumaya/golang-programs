package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acmpca"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := &aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			"AWS_ACCESS_KEY_ID",
			"AWS_SECRET_ACCESS_KEY",
			"AWS_SESSION_TOKEN"),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		log.Error(err)
	}

	// Check if credentials retrieved successfully
	_, err = sess.Config.Credentials.Get()
	if err != nil {
		log.Error(err)
	}

	svc := acmpca.New(sess)
	v, _ := svc.ListCertificateAuthorities(&acmpca.ListCertificateAuthoritiesInput{})
	fmt.Printf("%v", v)

}
