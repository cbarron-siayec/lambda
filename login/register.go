package main

import (
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(preSignUp events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignupResponse, error) {
	preSignUp.Response.AutoConfirmUser = false
	domain := strings.SplitAfter(preSignUp.Request.UserAttributes["email"], "@")
	clean := domain[1]
	ourDomain := "grupo-siayec.com.mx"
	if ourDomain == clean {
		preSignUp.Response.AutoConfirmUser = true
		log.Print("User: " + preSignUp.Request.UserAttributes["email"] + " confirmed")
		return preSignUp.Response, nil
	}
	log.Print("User: " + preSignUp.Request.UserAttributes["email"] + " was not confirmed")
	return preSignUp.Response, nil
}

func main() {
	lambda.Start(handler)
}
