package utils

import (
	"fmt"
	"time"
)

//CurrentDate returns the user current date
func CurrentDate() string {
	return time.Now().Format(time.RFC850)
}

//FloatToStr //converts a float time to string
func FloatToStr(num float64) string {
	return fmt.Sprintf("%f", num)
}
