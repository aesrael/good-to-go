package main

import (
	"good-to-go/utils"

	"github.com/fatih/color"
)

func main() {
	location := GetLocation()
	weather := GetWeatherInfo(location)
	printWeatherInfo(weather)
}

func printWeatherInfo(weather CurrentWeather) {
	color := color.New(color.FgRed).Add(color.Bold)

	color.Printf(weather.Currently.Summary + "\n")

	color.Printf("Location:           %s\n", weather.Timezone)
	color.Printf("Date:               %s\n", utils.CurrentDate())
	color.Printf("Temperature:        %f\n", weather.Currently.Temperature)
	color.Printf("Pressure:           %f\n", weather.Currently.Pressure)
	color.Printf("Visibility:         %f\n", weather.Currently.Visibility)
	color.Printf("WindSpeed:          %f\n", weather.Currently.WindSpeed)
	color.Printf("Cloud Cover:        %f\n", weather.Currently.CloudCover)

	color.Printf("\nSummary:            \n %s\n %s\n\n", weather.Hourly.Summary, weather.Daily.Summary)
}
