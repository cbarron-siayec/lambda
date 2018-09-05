package main

import (
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(preSignUp events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignup, error) {
	domain := strings.SplitAfter(preSignUp.Request.UserAttributes["email"], "@")
	clean := domain[1]
	ourDomain := "grupo-siayec.com.mx"
	if ourDomain == clean {
		log.Print("User: " + preSignUp.Request.UserAttributes["email"] + " created")
		return preSignUp, nil
	}
	workAround := preSignUp.Request.UserAttributes["produce-error"]
	preSignUp.Request.UserAttributes["username"] = workAround
	preSignUp.Request.UserAttributes["email"] = ""
	return preSignUp, nil
}

func main() {
	lambda.Start(handler)
}
