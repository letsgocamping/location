package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	port string
)

func init() {
	port = os.Getenv("LOCATION_SERVICE_PORT")
	if port == "" {
		port = ":8080"
	}
}

func main() {
	// Hello world, the web server

	http.HandleFunc("/api/location_service/midpoint", midpointHandler)
	http.ListenAndServe(":8080", nil)
}

type midpointBody struct {
	Cities [][]string `json:"cities"`
}

type midpointRes struct {
	Lat   float64 `json:"lat"`
	Long  float64 `json:"lon"`
	State string  `json:"state"`
	Parks string  `json:"parks"`
}

type errorRes struct {
	Error string `json:"error"`
}

func midpointHandler(w http.ResponseWriter, req *http.Request) {

	body, err := parseJson(req.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ll := getLatAndLong(body.Cities)

	lat, lng := getLatLngCenter(ll)

	state := findState(lat, lng)

	parks := getParks(state)

	res, err := json.Marshal(midpointRes{lat, lng, state, parks})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func parseJson(body io.ReadCloser) (midpointBody, error) {
	var request midpointBody
	buf, err := ioutil.ReadAll(body)
	defer body.Close()

	if err != nil {
		return midpointBody{}, err
	}

	err = json.Unmarshal(buf, &request)
	if err != nil {
		return midpointBody{}, err
	}

	return request, nil
}
