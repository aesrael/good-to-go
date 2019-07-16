package utils

import "log"

//HandleErr //handlerr
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
