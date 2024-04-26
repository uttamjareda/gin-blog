package main

import (
	"go-blog/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set the Gin mode
	gin.SetMode(gin.ReleaseMode) // Use gin.DebugMode during development

	// Initialize the router
	router := routes.SetupRouter()

	// Define the port to run the server on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if none specified
	}

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
