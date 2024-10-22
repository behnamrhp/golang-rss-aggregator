package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/behnamrhp/golang-rss-aggregator.git/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if params.Name == "" || params.Url == "" {
		respondWithError(w, 400, "Please provide correct url and name for you feed")
		return
	}

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create the feed: %v", err))
		return
	}

	respondWithJSON(w, 201, feedDbToModelDto(feed))
}

func (apiCfg *apiConfig) getFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't find the feeds: %v", err))
		return
	}
	respondWithJSON(w, 200, feedsDbToModelDto(feeds))
}
