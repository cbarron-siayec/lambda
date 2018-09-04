package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(preSignUp events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignupResponse, error) {
	domain := preSignUp.UserName
	log.Print(domain)
	return preSignUp.Response, nil
}

func main() {
	lambda.Start(handler)
}
