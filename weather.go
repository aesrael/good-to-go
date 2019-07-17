package main

import (
	"encoding/json"
	"fmt"
	"good-to-go/utils"
	"io/ioutil"
	"net/http"
)

//CurrentWeather struct
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

//GetWeatherInfo using dark sky api
func GetWeatherInfo(location Location) CurrentWeather {
	fmt.Println("\nfetching results...🌈")

	lat := utils.FloatToStr(location.Location.Lat)
	lng := utils.FloatToStr(location.Location.Lng)
	//get weather info using darksky weather api
	apiURL := "https://api.darksky.net/forecast/f55f3c1a9d0ccf6b5f88295b9a9aaf62/" + lat + "," + lng + "?exclude=flags,offset" + "&time=2019-02"
	// fmt.Println(apiURL)
	resp, err := http.Get(apiURL)
	utils.HandleErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var currentWeather CurrentWeather
	json.Unmarshal(body, &currentWeather)
	return currentWeather
}
