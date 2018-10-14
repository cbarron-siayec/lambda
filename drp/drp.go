package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/request"
)

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

func (c *sns) PublishRequest(input *PublishInput) (req *request.Request, output *PublishOutput) {
	req, resp := client.PublishRequest(params)
	err := req.Send()
	if err == nil { // resp is now filled
		log.Println(resp)
	}
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
	if blips.BLIPS > 0 {
		log.Print("System OK with blips:")
		log.Print(blips.BLIPS)
		return blips.BLIPS, nil
	}
	log.Print("System is Offline, admin warning ON")
	log.Print(blips.BLIPS)
	return blips.BLIPS, nil
}

func main() {
	lambda.Start(handler)
}
