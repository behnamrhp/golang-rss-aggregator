package main

import "net/http"

func handlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
