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

func handler(ctx context.Context, kinesisEvent KinesisAnalyticsEvent) (int, error) {
	sess := session.Must(session.NewSession())
	svc := sns.New(sess)
	paramsOK := &sns.PublishInput{
		Message:  aws.String("El servidor esta en linea, ESTATUS:OK"),
		TopicArn: aws.String("arn:aws:sns:us-east-1:890650648390:SERVER_HEALTH"),
	}
	paramsNotOK := &sns.PublishInput{
		Message:  aws.String("El servidor esta fuera de linea, si quiere activar el DRP acceda a la siguiente liga https://480lkm2mnb.execute-api.us-east-1.amazonaws.com/PRO/startInstance?code=@S!4y3c."),
		TopicArn: aws.String("arn:aws:sns:us-east-1:890650648390:SERVER_HEALTH"),
	}
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
		resp, err := svc.Publish(paramsOK)
		if err != nil {
			log.Print(err.Error())
		}
		log.Print(resp)
		return blips.BLIPS, nil
	}
	log.Print("System is Offline, admin warning ON")
	log.Print(blips.BLIPS)
	resp, err := svc.Publish(paramsNotOK)
	if err != nil {
		log.Print(err.Error())
	}
	log.Print(resp)
	return blips.BLIPS, nil
}

func main() {
	lambda.Start(handler)
}
