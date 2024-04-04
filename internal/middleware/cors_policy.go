package middleware

import (
	"net/http"
	"rcp-api-data/internal/config/security"
	"strings"

	"github.com/rs/cors"
)

type CorsMiddleware struct {
	cfg *security.Environment
}

func NewCorsMiddleware(cfg *security.Environment) *CorsMiddleware {
	return &CorsMiddleware{
		cfg: cfg,
	}
}

func (c *CorsMiddleware) Config(h http.Handler) http.Handler {
	hosts := c.cfg.CORS_ORIGIN
	allowedOrigins := strings.Split(hosts, ",")
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	return corsHandler.Handler(h)
}
