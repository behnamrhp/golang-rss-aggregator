package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/behnamrhp/golang-rss-aggregator.git/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

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

	connection, err := sql.Open("postgres", envs.dbUrl)

	if err != nil {
		log.Fatal("Can't connect to the database:", err)
	}

	apiConfig := apiConfig{
		DB: database.New(connection),
	}

	initRoutes(router, &apiConfig)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + envs.port,
	}

	fmt.Printf("Server is runnin on the port: %v", envs.port)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func initRoutes(router *chi.Mux, apiCfg *apiConfig) {
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerHealthCheck)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/users", apiCfg.createUser)
	v1Router.Get("/users", apiCfg.getUser)
	router.Mount("/v1", v1Router)
}
