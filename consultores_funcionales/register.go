package main

import (
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(preSignUp events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignup, error) {
	domain := strings.SplitAfter(preSignUp.Request.UserAttributes["username"], "@")
	preSignUp.Response["email"] = preSignUp.Request.UserAttributes["username"]
	clean := domain[1]
	ourDomain := "grupo-siayec.com.mx"
	if ourDomain == clean {
		log.Print("User: " + preSignUp.Request.UserAttributes["email"] + " created")
		return preSignUp, nil
	}
	produceError := domain[10]
	preSignUp.Request.UserAttributes["username"] = produceError
	return preSignUp, nil
}

func main() {
	lambda.Start(handler)
}
