package utils

import (
	"errors"
	"net/http"
	"rcp-api-data/internal/config"
	"rcp-api-data/internal/middleware"
)

// GetContextToken is a function that returns the token from the context

func GetContextToken(r *http.Request) (config.TokenFromCookie, error) {
	token := r.Context().Value(middleware.TokenKey)
	if token == nil {
		return config.TokenFromCookie{}, errors.New("token not found in context")
	}
	return token.(config.TokenFromCookie), nil
}