package main

import (
	"cnf-q/pkg/queueservice"
	"log"
	"os"
)

func main() {
	qs := queueservice.NewQueueService(
		queueservice.WithPort(os.Getenv("CNF_Q_PORT")),
		queueservice.WithAccessToken(os.Getenv("CNF_Q_ACCESS_TOKEN")),
		queueservice.WithTLS(
			os.Getenv("CNF_Q_TLS_CERT_FILE"),
			os.Getenv("CNF_Q_TLS_KEY_FILE"),
		))
	if err := qs.Run(); err != nil {
		log.Fatalf("error starting queue service: %v", err)
	}
}
