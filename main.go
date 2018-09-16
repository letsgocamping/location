package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(midpointHandler)
}

type midpointBody struct {
	Cities [][]string `json:"cities"`
}

type midpointRes struct {
	Headers         map[string]string `json:"headers"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	StatusCode      int               `json:"statusCode"`
	Body            body              `json:"body"`
}

type body struct {
	Lat    float64     `json:"lat"`
	Long   float64     `json:"lon"`
	State  string      `json:"state"`
	Parks  string      `json:"parks"`
	Cities [][]float64 `json:"cities"`
}

type errorRes struct {
	Error string `json:"error"`
}

func midpointHandler(req midpointBody) (events.APIGatewayProxyResponse, error) {
	headers := make(map[string]string)

	cities := getLatAndLong(req.Cities)

	lat, lng := getLatLngCenter(cities)
	headers["Location"] = fmt.Sprintf("%f, %f", lat, lng)

	state := findState(lat, lng)

	parks := getParks(state)

	Body := body{lat, lng, state, parks, cities}

	bodyJson, err := json.Marshal(Body)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{}, err
	}

	res := events.APIGatewayProxyResponse{Headers: headers, IsBase64Encoded: false, StatusCode: 200, Body: string(bodyJson)}

	return res, nil
}
