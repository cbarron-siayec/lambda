package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(kinesisEvent events.KinesisEvent) error {
	log.Print(kinesisEvent)
	log.Print(kinesisEvent.Records[0].EventSourceArn)
	return nil
}

func main() {
	lambda.Start(handler)
}
