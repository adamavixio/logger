package logger

import (
	"context"
	"log"
	"time"
)

func Trace(message string) {
	log.Printf("TRACE: %s", message)
}

func Info(message string) {
	log.Printf("INFO: %s", message)
}

func Warn(message string, err error) {
	if err != nil {
		log.Printf("WARNING: %s: %v", message, err)
	}
}

func Error(message string, err error) {
	if err != nil {
		log.Fatalf("ERROR - %s: %v", message, err)
	}
}

func Every(message string, interval time.Duration) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				log.Println(message)
				time.Sleep(interval)
			}
		}
	}()

	return cancel
}
