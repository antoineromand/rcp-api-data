package middleware

import (
	"net/http"
	"rcp-api-data/internal/config/security"
	"strings"

	"github.com/rs/cors"
)

func CorsMiddleware(h http.Handler, cfg *security.Environment) http.Handler {
	hosts := cfg.CORS_ORIGIN
	allowedOrigins := strings.Split(hosts, ",")
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	return corsHandler.Handler(h)
}
