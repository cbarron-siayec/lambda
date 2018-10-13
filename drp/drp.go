package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

type Blip struct {
	BLIP_COUNT string `json:"BLIP_COUNT"`
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
		log.Print("Not OK")
		return err.Error(), nil
	}
	res, err := strconv.ParseInt(blips.BLIP_COUNT, 10, 64)
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
