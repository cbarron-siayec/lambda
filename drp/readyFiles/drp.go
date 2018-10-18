package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sns"
)

type Blip struct {
	ID        string `json:"ID"`
	Author    string `json:"Author"`
	Timestamp string `json:"Timestamp"`
	Status    string `json:"Status"`
	Snapcount int    `json:"Snapcount"`
}

var sess = session.Must(session.NewSession())
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var db = dynamodb.New(sess, aws.NewConfig().WithRegion("us-east-1"))

func getItem(id string) (*Blip, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Server_Health"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
	}
	result, err := db.GetItem(input)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	blip := new(Blip)
	err = dynamodbattribute.UnmarshalMap(result.Item, &blip)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return blip, nil
}

func putItem(nuevoRegistro *Blip) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Server_Health"),
		Item: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(nuevoRegistro.ID),
			},
			"Author": {
				S: aws.String(nuevoRegistro.Author),
			},
			"Timestamp": {
				S: aws.String(nuevoRegistro.Timestamp),
			},
			"Status": {
				S: aws.String(nuevoRegistro.Status),
			},
			"Snapcount": {
				N: aws.String(strconv.Itoa(nuevoRegistro.Snapcount)),
			},
		},
	}
	result, err := db.PutItem(input)
	if err != nil {
		log.Print(err)
		log.Print(result)
	}
	return err
}

func handler(ctx context.Context) (int, error) {
	svc := sns.New(sess)
	paramsNotOK := &sns.PublishInput{
		Message:  aws.String("Servidor offline, codigo: @S!4y3c. https://s3.amazonaws.com/gsiayec-drp-start/index.html"),
		TopicArn: aws.String("arn:aws:sns:us-east-1:890650648390:SERVER_HEALTH"),
	}
	paramsCritical := &sns.PublishInput{
		Message:  aws.String("DRP, esta por comenzar automáticamente si no toma una acción"),
		TopicArn: aws.String("arn:aws:sns:us-east-1:890650648390:SERVER_HEALTH"),
	}

	blip, err := getItem("D4m0")
	if err != nil {
		log.Print(err)
		return 200, nil
	}

	switch blip.Snapcount {
	case 1:
		putItem(&Blip{
			ID:        "D4m0",
			Author:    "AMZ",
			Timestamp: time.Now().UTC().String(),
			Status:    "OK",
			Snapcount: 0,
		})
		log.Print("System OK with blips:" + strconv.Itoa(blip.Snapcount))
		return blip.Snapcount, nil
	case 0:
		putItem(&Blip{
			ID:        "D4m0",
			Author:    "AMZ",
			Timestamp: time.Now().UTC().String(),
			Status:    "Alert 1",
			Snapcount: -3,
		})
		log.Print("System is Offline, admin warning ON Snapcount is:" + strconv.Itoa(blip.Snapcount))
		resp, err := svc.Publish(paramsNotOK)
		if err != nil {
			log.Print(err)
		}
		log.Print(resp.String())
		return blip.Snapcount, nil
	case -1:
		putItem(&Blip{
			ID:        "D4m0",
			Author:    "AMZ",
			Timestamp: time.Now().UTC().String(),
			Status:    "Alert 1",
			Snapcount: -3,
		})
		log.Print("System is Offline, admin warning ON Snapcount is:" + strconv.Itoa(blip.Snapcount))
		resp, err := svc.Publish(paramsNotOK)
		if err != nil {
			log.Print(err)
		}
		log.Print(resp.String())
		return blip.Snapcount, nil
	case -2:
		putItem(&Blip{
			ID:        "D4m0",
			Author:    "AMZ",
			Timestamp: time.Now().UTC().String(),
			Status:    "Alert 2",
			Snapcount: -3,
		})
		log.Print("System is Offline, admin warning ON Snapcount is:" + strconv.Itoa(blip.Snapcount))
		resp, err := svc.Publish(paramsNotOK)
		if err != nil {
			log.Print(err)
		}
		log.Print(resp.String())
		return blip.Snapcount, nil
	case -3:
		putItem(&Blip{
			ID:        "D4m0",
			Author:    "AMZ",
			Timestamp: time.Now().UTC().String(),
			Status:    "Alert 3",
			Snapcount: -5,
		})
		log.Print("System is Offline, DRP will be implemented now:" + strconv.Itoa(blip.Snapcount))
		resp, err := svc.Publish(paramsCritical)
		if err != nil {
			log.Print(err)
		}
		log.Print(resp.String())
		svc := ec2.New(sess)
		paramsEC2 := &ec2.StartInstancesInput{
			InstanceIds: []*string{aws.String("i-0720a0a34dc424244")},
		}
		callbackEC2, err := svc.StartInstances(paramsEC2)
		if err != nil {
			log.Print(err)
			return -15, nil
		}
		log.Print(callbackEC2)
		return blip.Snapcount, nil
	default:
		putItem(&Blip{
			ID:        "D4m0",
			Author:    "AMZ",
			Timestamp: time.Now().UTC().String(),
			Status:    "Monitor Error or DRP Started",
			Snapcount: -30,
		})
	}

	return -45, nil
}

func main() {
	lambda.Start(handler)
}
