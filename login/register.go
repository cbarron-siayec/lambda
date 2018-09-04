package main

import (
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(preSignUp events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignupResponse, error) {
	preSignUp.Response.AutoConfirmUser = false
	log.Print("Logs Start Here")
	domain := strings.SplitAfter(preSignUp.Request.UserAttributes["email"], "@")
	log.Print(preSignUp.Request.UserAttributes["email"])
	clean := domain[1]
	log.Print(clean)
	ourDomain := "grupo-siayec.com.mx"
	if ourDomain == clean {
		preSignUp.Response.AutoConfirmUser = true
		return preSignUp.Response, nil
	}
	return preSignUp.Response, nil
}

func main() {
	lambda.Start(handler)
}
