package main

import (
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(preSignUp events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignup, error) {
	log.Print(preSignUp.UserName)
	domain := strings.SplitAfter(preSignUp.UserName, "@")
	log.Print(domain)
	preSignUp.Request.UserAttributes["email"] = preSignUp.UserName
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
