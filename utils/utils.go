package utils

import (
	"fmt"
	"math"
	"time"
)

//CurrentDate returns the user current date
func CurrentDate() string {
	return time.Now().Format(time.RFC850)
}

//FloatToStr //converts a float type to string
func FloatToStr(num float64) string {
	return fmt.Sprintf("%f", num)
}

func RoundNum(num float64) float64 {
	return math.Round((num * 100) / 100)
}
