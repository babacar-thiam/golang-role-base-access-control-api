package main

import (
	"log"
	"role-base-access-control-api/configs"
	"role-base-access-control-api/internal/app"
)

func main() {
	// Load the application configuration
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Create a new application instance
	application, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("Error creating application: %v", err)
	}

	// Run the application
	log.Printf("Starting the application...")
	application.RunApp()
}
