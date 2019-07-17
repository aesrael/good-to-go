package main

import (
	"encoding/json"
	"good-to-go/utils"
	"io/ioutil"
	"net/http"
)

// Location struct
type Location struct {
	IP       string `json:"ip"`
	Location struct {
		Country    string  `json:"country"`
		Region     string  `json:"region"`
		City       string  `json:"city"`
		Lat        float64 `json:"lat"`
		Lng        float64 `json:"lng"`
		PostalCode string  `json:"postalCode"`
		Timezone   string  `json:"timezone"`
		GeonameID  int     `json:"geonameId"`
	} `json:"location"`
}

//GetLocation location from ipaddress using ipify
func GetLocation() Location {
	// https://geo.ipify.org/api/v1?apiKey=[API_KEY]
	resp, err := http.Get("https://geo.ipify.org/api/v1?apiKey=at_yBQonkqZ1uamplShTvYIbodjpCpJ9")
	utils.HandleErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	utils.HandleErr(err)
	var location Location
	json.Unmarshal(body, &location)
	return location
}
