package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	configs := getEnvs()

	runWebServer(configs)
}

func runWebServer(envs Envs) {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	initRoutes(router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + envs.port,
	}

	fmt.Printf("Server is runnin on the port: %v", envs.port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func initRoutes(router *chi.Mux) {
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerHealthCheck)
	v1Router.Get("/err", handlerError)
	router.Mount("/v1", v1Router)
}
