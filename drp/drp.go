package main

import (
	"context"
	"encoding/base64"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

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

func handler(ctx context.Context, kinesisEvent KinesisAnalyticsEvent) (string, error) {
	encoded := kinesisEvent.Record[0].Data
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	log.Print("DATA: " + string(decoded))
	res, err := strconv.ParseInt(string(decoded), 10, 64)
	if err != nil {
		log.Print("Not OK")
		return "Not OK", nil
	}
	if res > 0 {
		log.Print(string(res))
		return string(res), nil
	}
	return string(decoded), nil
}

func main() {
	lambda.Start(handler)
}
