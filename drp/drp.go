package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-lambda-go/events"
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

type Blip struct {
	BLIPS int `json:"BLIPS"`
}

type Records struct {
	RecordId string `json:"recordId"`
	Data     string `json:"data"`
}

type KinesisAnalyticsEvent struct {
	InvocationId   string    `json:"invocationId"`
	ApplicationArn string    `json:"applicationArn"`
	StreamArn      string    `json:"streamArn"`
	Record         []Records `json:"records"`
}

func handler(ctx context.Context, kinesisEvent KinesisAnalyticsEvent, apiEvent events.APIGatewayProxyRequest) (int, events.APIGatewayProxyResponse, error) {
	if apiEvent.HTTPMethod = "POST" {
		log.Print(apiEvent.QueryStringParameters["code"])
		log.Print("POSTED!")
		return -1,events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    map[string]string{"Headers": `[Content-Type:application/json]`},
			Body:       "DRP Initiated",
		} ,nil
	}
	sess := session.Must(session.NewSession())
	svc := sns.New(sess)
	paramsOK := &sns.PublishInput{
		Message:  aws.String("OK"),
		TopicArn: aws.String("arn:aws:sns:us-east-1:890650648390:SERVER_HEALTH"),
	}
	paramsNotOK := &sns.PublishInput{
		Message:  aws.String("El servidor esta fuera de linea, si quiere activar el DRP acceda a la siguiente liga http://www.google.com"),
		TopicArn: aws.String("arn:aws:sns:us-east-1:890650648390:SERVER_HEALTH"),
	}
	encoded := kinesisEvent.Record[0].Data
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	blips := new(Blip)
	err := json.Unmarshal([]byte(decoded), &blips)
	if err != nil {
		log.Print(err.Error())
		return -1,events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    map[string]string{"Headers": `[Content-Type:application/json]`},
			Body:       string("Error converting Kinesis"),
		} ,nil
	}
	if blips.BLIPS > 0 {
		log.Print("System OK with blips:")
		log.Print(blips.BLIPS)
		resp, err := svc.Publish(paramsOK)
		if err != nil {
			log.Print(err.Error())
		}
		log.Print(resp)
		return blips.BLIPS, events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    map[string]string{"Headers": `[Content-Type:application/json]`},
			Body:       string("Machine Up"),
		} ,nil
	}
	log.Print("System is Offline, admin warning ON")
	log.Print(blips.BLIPS)
	resp, err := svc.Publish(paramsNotOK)
	if err != nil {
		log.Print(err.Error())
	}
	log.Print(resp)
	return blips.BLIPS, events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Headers": `[Content-Type:application/json]`},
		Body:       string("System is Offline, admin warning ON"),
	} ,nil
}

func main() {
	lambda.Start(handler)
}
