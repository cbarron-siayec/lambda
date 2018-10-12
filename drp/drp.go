package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

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
	log.Print(kinesisEvent.ApplicationArn)
	log.Print(kinesisEvent.StreamArn)
	return fmt.Sprintf("Data BASE64: " + kinesisEvent.ApplicationArn), nil
}

func main() {
	lambda.Start(handler)
}
