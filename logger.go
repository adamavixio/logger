package logger

import "log"

func HandleError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
