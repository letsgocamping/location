package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getParks(state string) string {

	url := fmt.Sprintf("https://api.nps.gov/api/v1/parks?stateCode=%s&limit=50&q=national%spark", state, "%20")

	req, err := http.Get(url)
	if err != nil {

	}

	defer req.Body.Close()

	v, _ := ioutil.ReadAll(req.Body)

	return string(v)
}
