package main

import (
	"fmt"
	"net/http"

	"github.com/behnamrhp/golang-rss-aggregator.git/internal/auth"
	"github.com/behnamrhp/golang-rss-aggregator.git/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Auth Error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't find the user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
