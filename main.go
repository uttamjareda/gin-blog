package main

import (
	"go-blog/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the Gin mode
	gin.SetMode(gin.ReleaseMode) // Use gin.DebugMode during development

	// Initialize the router
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // or "*" for allowing any origin
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000" // Adjust the allowed origin here
		},
		MaxAge: 12 * time.Hour,
	}))

	// Initialize the router
	routes.SetupRouter(router)

	// Define the port to run the server on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if none specified
	}

	// Start the server
	log.Printf("Starting server on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
