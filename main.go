package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	port := getPort()

	runWebServer(port)
}

func getPort() string {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port is not set")
	}
	return port
}

func runWebServer(port string) {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	fmt.Println("here hereeeee")
	v1Router.Get("/healthz", handlerHealthCheck)
	router.Mount("/v1", v1Router)
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
