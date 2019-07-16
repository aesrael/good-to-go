package main

import (
	"encoding/json"
	"fmt"
	"good-to-go/utils"
	"io/ioutil"
	"net/http"
)

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

type CurrentWeather struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Currently struct {
		Time                int     `json:"time"`
		Summary             string  `json:"summary"`
		Icon                string  `json:"icon"`
		PrecipIntensity     int     `json:"precipIntensity"`
		PrecipProbability   int     `json:"precipProbability"`
		Temperature         float64 `json:"temperature"`
		ApparentTemperature float64 `json:"apparentTemperature"`
		DewPoint            float64 `json:"dewPoint"`
		Humidity            float64 `json:"humidity"`
		Pressure            float64 `json:"pressure"`
		WindSpeed           float64 `json:"windSpeed"`
		WindGust            float64 `json:"windGust"`
		WindBearing         int     `json:"windBearing"`
		CloudCover          float64 `json:"cloudCover"`
		UvIndex             int     `json:"uvIndex"`
		Visibility          float64 `json:"visibility"`
		Ozone               float64 `json:"ozone"`
	} `json:"currently"`
	Hourly struct {
		Summary string `json:"summary"`
	} `json:"hourly"`
	Daily struct {
		Summary string `json:"summary"`
	} `json:"daily"`
}

func main() {
	location := getLocationFromIP()
	weather := getWeatherInfo(location)
	// fmt.Println(weather.Daily, "summary")
	// printWeatherInfo(weather)
	fmt.Printf("%+v\n", weather)
}

// func printWeatherInfo(weather) CurrentWeather {

// }
func getWeatherInfo(location Location) CurrentWeather {
	lat := floatToStr(location.Location.Lat)
	lng := floatToStr(location.Location.Lng)
	//get weather info using darksky weather api
	apiURL := "https://api.darksky.net/forecast/f55f3c1a9d0ccf6b5f88295b9a9aaf62/" + lat + "," + lng + "?exclude=flags,offset" + "&time=2019-02"
	fmt.Println(apiURL)
	resp, err := http.Get(apiURL)
	utils.HandleErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var currentWeather CurrentWeather
	json.Unmarshal(body, &currentWeather)
	return currentWeather
}

func floatToStr(num float64) string {
	return fmt.Sprintf("%f", num)
}

//get user location from ipaddress using ipify
func getLocationFromIP() Location {
	resp, err := http.Get("https://geo.ipify.org/api/v1?apiKey=at_yBQonkqZ1uamplShTvYIbodjpCpJ9")
	utils.HandleErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	utils.HandleErr(err)
	var location Location
	json.Unmarshal(body, &location)
	return location
}
