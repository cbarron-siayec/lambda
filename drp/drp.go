package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type Blip struct {
	BLIP_COUNT int `json:"BLIP_COUNT"`
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

func handler(ctx context.Context, kinesisEvent KinesisAnalyticsEvent) (string, error) {
	encoded := kinesisEvent.Record[0].Data
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	blips := new(Blip)
	err := json.Unmarshal([]byte(decoded), &blips)
	if err != nil {
		log.Print(err.Error())
		return err.Error(), nil
	}
	log.Print(blips.BLIP_COUNT)
	res := blips.BLIP_COUNT
	if blips.BLIP_COUNT > 0 {
		log.Print("On If" + string(blips.BLIP_COUNT))
		return string(res), nil
	}
	log.Print("Out of if" + string(blips.BLIP_COUNT))
	return string(blips.BLIP_COUNT), nil
}

func main() {
	lambda.Start(handler)
}
