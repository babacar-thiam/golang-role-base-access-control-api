// cmd/main.go
package main

import (
	"log"
	"role-base-access-control-api/configs"
	"role-base-access-control-api/internal/app"
)

// @title Role-Based Access Control API
// @version 1.0
// @description API for managing roles and users with RBAC.
// @host localhost:3001
// @BasePath /api/v1
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
