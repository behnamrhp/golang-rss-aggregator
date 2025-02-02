package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

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

	db := connectDatabase(configs.dbUrl)

	go startScraping(db, 10, time.Minute)

	runWebServer(configs.port, db)
}

func connectDatabase(dbUrl string) *database.Queries {
	connection, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Can't connect to the database:", err)
	}

	return database.New(connection)
}

func runWebServer(port string, db *database.Queries) {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	apiConfig := apiConfig{
		DB: db,
	}

	initRoutes(router, &apiConfig)

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

func initRoutes(router *chi.Mux, apiCfg *apiConfig) {
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerHealthCheck)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/users", apiCfg.createUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.getUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.createFeed))
	v1Router.Get("/feeds", apiCfg.getFeeds)
	v1Router.Post("/feed-follows", apiCfg.middlewareAuth(apiCfg.createFeedFollow))
	v1Router.Get("/feed-follows", apiCfg.middlewareAuth(apiCfg.getFeedFollows))
	v1Router.Delete("/feed-follows/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.deleteFeedFollow))
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.getPostsForUser))
	router.Mount("/v1", v1Router)
}
