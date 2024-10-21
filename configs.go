package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Envs struct {
	port  string
	dbUrl string
}

func getEnvs() Envs {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port is not set")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB_URL is not set")
	}
	return Envs{
		port:  port,
		dbUrl: dbUrl,
	}
}
