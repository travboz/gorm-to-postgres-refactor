package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/travboz/go-quest/controllers"
	"github.com/travboz/go-quest/models"
)

func main() {
	godotenv.Load()

	routes := controllers.New()
	models.ConnectDatabase()

	server := &http.Server{
		Addr:         "0.0.0.0:8000",
		Handler:      routes,
		ReadTimeout:  10 * time.Second, // Prevent slow read attacks
		WriteTimeout: 10 * time.Second, // Prevent slow write attacks
		IdleTimeout:  60 * time.Second, // Close idle connections after a duration

	}

	// Log server start message
	log.Println("Server is running on port 8000")

	// Start server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
