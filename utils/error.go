package utils

import "log"

//HandleErr //generic error handler
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
