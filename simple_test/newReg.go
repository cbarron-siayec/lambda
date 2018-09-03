package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-west-1"))

type NewReg struct {
	Name string `json:"Name"`
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if req.Headers["Content-Type"] != "application/json" {
		return clientError(http.StatusMovedPermanently)
	}

	nuevoRegistro := new(NewReg)
	err := json.Unmarshal([]byte(req.Body), &nuevoRegistro)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}
	if nuevoRegistro.Name == "" {
		return clientError(http.StatusBadRequest)
	}

	err = putItem(nuevoRegistro)
	if err != nil {
		return serverError(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Registro": fmt.Sprintln(req), "Access-Control-Allow-Origin": "*"},
	}, nil
}

func putItem(nuevoRegistro *NewReg) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Names"),
		Item: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(nuevoRegistro.Name),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}

func main() {
	lambda.Start(handler)
}
