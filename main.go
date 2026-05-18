package main

import (
	"project-management/config"
	"project-management/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	// Connect PostgreSQL Database
	config.ConnectDB()

	// Create Gin router
	r := gin.Default()

	r.Use(cors.Default())

	// Setup API routes
	routes.SetupRoutes(r)

	

	// Run server
	r.Run(":8081")
}
