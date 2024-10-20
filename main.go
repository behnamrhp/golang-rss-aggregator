package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	port := getPort()

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Server is runnin on the port: %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func getPort() string {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port is not set")
	}
	return port
}
