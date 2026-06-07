package main

import (
	"log"

	"erp-project-management/config"
	"erp-project-management/database"
	"erp-project-management/routes"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database and run migrations
	database.Connect(cfg)

	// Seed the database with initial data
	database.Seed()

	// Setup router with all routes and middleware
	r := routes.SetupRouter(cfg)

	// Start server
	addr := ":" + cfg.ServerPort
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
