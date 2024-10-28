package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Fenroe/carbonarapi/internal/config"
	"github.com/Fenroe/carbonarapi/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Export variables in .env file
	godotenv.Load()
	// Get port and db url
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	// Connect to database. If it fails then quit and log the error
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	// Initialize config struct
	apiConfig := config.Config{
		Greeting: "Hi Banana!",
		Queries:  database.New(db),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, apiConfig.Greeting)
	})
	// Initialize server
	server := http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}
	// Feedback log
	fmt.Printf("Starting server on %s\n", server.Addr)
	// If server encounters an error then quit and log it
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
