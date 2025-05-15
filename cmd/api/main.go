package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/travboz/go-quest/internal/env"
)

func main() {
	godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := ConnectToDB(dsn)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Defer a call to db.Close() so that the connection pool is closed before the
	// main() function exits.
	defer db.Close()

	env := env.NewEnv(db)

	routes := routes(env)

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
