package main

import (
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(preSignUp events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignupResponse, error) {
	preSignUp.Response.AutoConfirmUser = false
	logs := preSignUp.UserName
	log.Print(logs)
	log.Print("This should be logged")
	domain := strings.SplitAfter(preSignUp.UserName, "@")[1]
	ourDomain := "grupo-siayec.com.mx"
	if ourDomain == domain {
		preSignUp.Response.AutoConfirmUser = true
		return preSignUp.Response, nil
	}
	return preSignUp.Response, nil
}

func main() {
	lambda.Start(handler)
}
