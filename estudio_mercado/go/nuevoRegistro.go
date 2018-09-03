package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-west-1"))

type estudio_mercado struct {
	Administracion        string `json:"administracion"`
	AguaDrenaje           string `json:"aguaDrenaje"`
	Alumbrado             string `json:"alumbrado"`
	CallesParquesJardines string `json:"callesParquesJardines"`
	ColaboradoresAcceso   string `json:"colaboradoresAcceso"`
	Comentarios           string `json:"comentarios"`
	DuplicarTrabajo       string `json:"duplicarTrabajo"`
	Fecha                 string `json:"fecha"`
	HorasDia              string `json:"horasDia"`
	InfraestructuraNube   string `json:"infraestructuraNube"`
	ManejoDatos           string `json:"manejoDatos"`
	Mercados              string `json:"mercados"`
	NombreColaborador     string `json:"nombreColaborador"`
	PersonalExclusivo     string `json:"personalExclusivo"`
	Presupuesto           string `json:"presupuesto"`
	SeguridadPublica      string `json:"seguridadPublica"`
	ServiciosLimpia       string `json:"serviciosLimpia"`
	Tramites              string `json:"tramites"`
	Utilidad              string `json:"utilidad"`
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	nuevoRegistro := new(estudio_mercado)
	err := json.Unmarshal([]byte(req.Body), &nuevoRegistro)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}
	if nuevoRegistro.Fecha == "" || nuevoRegistro.NombreColaborador == "" {
		return clientError(http.StatusBadRequest)
	}

	err = putItem(nuevoRegistro)
	if err != nil {
		return serverError(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Registro": fmt.Sprintln(req)},
	}, nil
}

func putItem(nuevoRegistro *estudio_mercado) error {

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Estudio_Mercado"),
		Item: map[string]*dynamodb.AttributeValue{
			"administracion": {
				S: aws.String(nuevoRegistro.Administracion),
			},
			"aguaDrenaje": {
				S: aws.String(nuevoRegistro.AguaDrenaje),
			},
			"alumbrado": {
				S: aws.String(nuevoRegistro.Alumbrado),
			},
			"callesParquesJardines": {
				S: aws.String(nuevoRegistro.CallesParquesJardines),
			},
			"colaboradoresAcceso": {
				S: aws.String(nuevoRegistro.ColaboradoresAcceso),
			},
			"comentarios": {
				S: aws.String(nuevoRegistro.Comentarios),
			},
			"duplicarTrabajo": {
				S: aws.String(nuevoRegistro.DuplicarTrabajo),
			},
			"fecha": {
				S: aws.String(nuevoRegistro.Fecha),
			},
			"horasDia": {
				S: aws.String(nuevoRegistro.HorasDia),
			},
			"infraestructuraNube": {
				S: aws.String(nuevoRegistro.InfraestructuraNube),
			},
			"mercados": {
				S: aws.String(nuevoRegistro.Mercados),
			},
			"nombreColaborador": {
				S: aws.String(nuevoRegistro.NombreColaborador),
			},
			"personalExclusivo": {
				S: aws.String(nuevoRegistro.PersonalExclusivo),
			},
			"presupuesto": {
				S: aws.String(nuevoRegistro.Presupuesto),
			},
			"seguridadPublica": {
				S: aws.String(nuevoRegistro.SeguridadPublica),
			},
			"serviciosLimpia": {
				S: aws.String(nuevoRegistro.ServiciosLimpia),
			},
			"tramites": {
				S: aws.String(nuevoRegistro.Tramites),
			},
			"utilidad": {
				S: aws.String(nuevoRegistro.Utilidad),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}

func main() {
	lambda.Start(handler)
}
