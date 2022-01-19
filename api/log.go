package api

import (
	"log"
	"os"
)

func logRequest(message string) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Panicf("Cannot lookup hostname: %s\n", err.Error())
	}

	log.Printf("[%s] req: %s\n", hostname, message)
}
