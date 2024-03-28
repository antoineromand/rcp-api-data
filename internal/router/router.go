package router

import (
	"net/http"
	"rcp-api-data/internal/config/security"
	controller_account "rcp-api-data/internal/controller/account"
	"rcp-api-data/internal/middleware"

	"gorm.io/gorm"
)

func Router(db *gorm.DB, cfg *security.Environment) {
	http.Handle("/api/account/me", middleware.CorsMiddleware(middleware.ValidateTokenMiddleware(http.HandlerFunc(controller_account.AccountController(db)), cfg)), cfg)
}
