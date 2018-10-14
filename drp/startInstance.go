package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

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

func handler(apiEvent events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if apiEvent.HTTPMethod == "POST" && apiEvent.QueryStringParameters["code"] == "@S!4y3c." {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    map[string]string{"Headers": `[Content-Type:application/json]`},
			Body:       "DRP Initiated",
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Headers": `[Content-Type:application/json]`},
		Body:       "Wrong Code or HTTP Method!",
	}, nil
}

func main() {
	lambda.Start(handler)
}
