package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/behnamrhp/golang-rss-aggregator.git/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedId,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create the feed: %v", err))
		return
	}

	respondWithJSON(w, 201, feedFollowDbToModelDto(feedFollow))
}

func (apiConfig *apiConfig) getFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	dbFeedFollows, err := apiConfig.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't find any feed follows: %v", err))
		return
	}
	respondWithJSON(w, 200, feedFollowsDbToModelDto(dbFeedFollows))
}
