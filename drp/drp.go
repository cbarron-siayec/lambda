package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, kinesisEvent events.KinesisEvent) error {
	log.Print(kinesisEvent)
	log.Print(kinesisEvent.Records[0].Kinesis.ApproximateArrivalTimestamp)
	log.Print(kinesisEvent.Records[0].Kinesis.EncryptionType)
	log.Print(kinesisEvent.Records[0].Kinesis.KinesisSchemaVersion)
	log.Print(kinesisEvent.Records[0].Kinesis.PartitionKey)
	log.Print(kinesisEvent.Records[0].Kinesis.SequenceNumber)
	log.Print(kinesisEvent.Records[0].Kinesis.Data)
	for _, record := range kinesisEvent.Records {
		kinesisRecord := record.Kinesis
		dataBytes := kinesisRecord.Data
		dataText := string(dataBytes)
		log.Print(record.Kinesis.Data)
		log.Print(dataText)

	}
	return nil
}

func main() {
	lambda.Start(handler)
}
