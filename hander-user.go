package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/behnamrhp/golang-rss-aggregator.git/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		FullName string `json:"full_name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if params.FullName == "" {
		respondWithError(w, 400, "Please provide correct full name as full_name")
		return
	}

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Fullname:  params.FullName,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 201, userDbToModelDto(user))
}

func (apiCfg *apiConfig) getUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, userDbToModelDto(user))
}
