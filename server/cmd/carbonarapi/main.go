package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Fenroe/carbonarapi/internal/config"
	"github.com/Fenroe/carbonarapi/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Export variables in .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Ignore this error if in production: %v\n", err)
	}
	// Get port and db url
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env variable not set")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL env variable not set")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET env variable not set")
	}
	// Connect to database. If it fails then quit and log the error
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	// Initialize config struct
	apiConfig := config.Config{
		Greeting:  "Hi Banana!",
		DB:        database.New(db),
		JWTSecret: jwtSecret,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, apiConfig.Greeting)
	})
	// Handlers
	mux.HandleFunc("POST /api/users", apiConfig.HandlerCreateUser)
	mux.HandleFunc("POST /api/register", apiConfig.HandlerCreateUser)
	mux.HandleFunc("POST /api/login", apiConfig.HandlerLogin)
	mux.HandleFunc("POST /api/refresh", apiConfig.HandlerRefresh)
	mux.HandleFunc("POST /api/revoke", apiConfig.HandlerRevoke)
	mux.HandleFunc("POST /api/clips", apiConfig.HandlerCreateClip)
	mux.HandleFunc("GET /api/clips", apiConfig.HandlerGetClipsByUser)

	// Initialize server
	server := http.Server{
		Handler:           mux,
		Addr:              ":" + port,
		ReadHeaderTimeout: time.Duration(5) * time.Second,
	}
	// Feedback log
	fmt.Printf("Starting server on %s\n", server.Addr)
	// If server encounters an error then quit and log it
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
