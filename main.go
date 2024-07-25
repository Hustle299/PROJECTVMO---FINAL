package main

import (
	"log"

	"github.com/cesc1802/onboarding-and-volunteer-service/cmd/config"
	"github.com/cesc1802/onboarding-and-volunteer-service/cmd/server"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the server
	srv := server.NewServer(cfg)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
