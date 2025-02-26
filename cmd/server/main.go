package main

import (
	"cnf-q/pkg/queueservice"
	"log"
)

func main() {
	qs := queueservice.NewQueueService()
	if err := qs.Run(); err != nil {
		log.Fatalf("error starting queue service: %v", err)
	}
}
