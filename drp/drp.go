package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(kinesisEvent events.KinesisRecord) error {
	log.Print(kinesisEvent)
	return nil
}

func main() {
	lambda.Start(handler)
}
