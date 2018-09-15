package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	apiKey string
)

func init() {
	apiKey = os.Getenv("API_KEY")
}

func getLatAndLong(cities [][]string) [][]float64 {
	var latLong [][]float64
	var wg sync.WaitGroup

	wg.Add(len(cities))

	for _, v := range cities {
		go func(v []string) {
			city := v[0]
			state := v[1]

			res, _ := http.Get(formatUrl(city, state))

			j, _ := parsePlacesJson(res.Body)

			lat := j.Results[0].Geometry.Location.Lat
			lng := j.Results[0].Geometry.Location.Long

			latLong = append(latLong, []float64{lat, lng})
			wg.Done()
		}(v)
	}

	wg.Wait()
	return latLong
}

func formatUrl(city, state string) string {
	city = strings.Replace(city, " ", "%20", -1)
	return fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s,%s&key=%s", city, state, apiKey)
}

type placesRes struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat  float64 `json:"lat"`
				Long float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

func parsePlacesJson(body io.ReadCloser) (placesRes, error) {
	var request placesRes
	buf, err := ioutil.ReadAll(body)
	defer body.Close()

	if err != nil {
		return placesRes{}, err
	}

	err = json.Unmarshal(buf, &request)
	if err != nil {
		return placesRes{}, err
	}

	return request, nil
}

type stateRes struct {
	plusCode struct {
		compoundCode string `json:"compound_code"`
	}
}

func findState(lat, lng float64) string {

	res, _ := http.Get(formatStateUrl(lat, lng))

}

func formatStateUrl(lat, lng float64) string {

	return fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%f,%f&key=%s", lat, lng, apiKey)

}
