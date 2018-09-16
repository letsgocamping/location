package main

import (
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

func midpointHandler(req midpointBody) (body, error) {
	cities := getLatAndLong(req.Cities)

	lat, lng := getLatLngCenter(cities)

	state := findState(lat, lng)

	parks := getParks(state)

	res := body{lat, lng, state, parks, cities}

	return res, nil
}
