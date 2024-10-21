package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	authErrorMessage := "Not Autherized"
	if val == "" {
		return "", errors.New(authErrorMessage)
	}

	splitedToken := strings.Split(val, " ")

	if len(splitedToken) != 2 {
		return "", errors.New(authErrorMessage)
	}

	if splitedToken[0] != "ApiKey" {
		return "", errors.New(authErrorMessage)
	}

	return splitedToken[1], nil
}
