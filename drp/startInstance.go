package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
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
	sess := session.Must(session.NewSession())
	svc := ec2.New(sess)
	params := &ec2.StartInstancesInput{
		InstanceIds: []*string{aws.String("i-0720a0a34dc424244")},
	}
	if apiEvent.HTTPMethod == "POST" && apiEvent.QueryStringParameters["code"] == "@S!4y3c." {
		resp, err := svc.StartInstances(params)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    map[string]string{"Headers": `[Content-Type:application/json]`},
				Body:       resp.StartingInstances[0].CurrentState.String(),
			}, nil
		}
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
