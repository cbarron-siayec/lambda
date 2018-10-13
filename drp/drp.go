package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type Blip struct {
	NUMBER_OF_DISTINCT_ITEMS int `json:"NUMBER_OF_DISTINCT_ITEMS"`
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

func handler(ctx context.Context, kinesisEvent KinesisAnalyticsEvent) (int, error) {
	encoded := kinesisEvent.Record[0].Data
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	blips := new(Blip)
	err := json.Unmarshal([]byte(decoded), &blips)
	if err != nil {
		log.Print(err.Error())
		return -1, nil
	}
	if blips.NUMBER_OF_DISTINCT_ITEMS > 0 {
		log.Print("System OK with blips:")
		log.Print(blips.NUMBER_OF_DISTINCT_ITEMS)
		return blips.NUMBER_OF_DISTINCT_ITEMS, nil
	}
	log.Print("System is Offline, admin warning ON")
	log.Print(blips.NUMBER_OF_DISTINCT_ITEMS)
	return blips.NUMBER_OF_DISTINCT_ITEMS, nil
}

func main() {
	lambda.Start(handler)
}
