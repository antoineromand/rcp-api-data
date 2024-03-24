package router

import (
	"net/http"
	"rcp-api-data/internal/config/security"
	controller_account "rcp-api-data/internal/controller/account"
	"rcp-api-data/internal/middleware"

	"gorm.io/gorm"
)

func Router(db *gorm.DB, cfg *security.Environment) {
    // http.Handle("/api/data", middleware.ValidateTokenMiddleware(http.HandlerFunc(controller.DataController())))
	// http.Handle("/api/data/insert", middleware.ValidateTokenMiddleware(http.HandlerFunc(controller.())))
	http.Handle("/api/account", middleware.ValidateTokenMiddleware(http.HandlerFunc(controller_account.AccountController(db)), cfg))
}