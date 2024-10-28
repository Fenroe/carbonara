package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "8080"
	server := http.Server{
		Handler: http.NewServeMux(),
		Addr:    ":" + port,
	}
	fmt.Printf("Starting server on %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
