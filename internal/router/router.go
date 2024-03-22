package router

import (
	"net/http"
	controller_account "rcp-api-data/internal/controller/account"
	"rcp-api-data/internal/middleware"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {
    // http.Handle("/api/data", middleware.ValidateTokenMiddleware(http.HandlerFunc(controller.DataController())))
	// http.Handle("/api/data/insert", middleware.ValidateTokenMiddleware(http.HandlerFunc(controller.())))
	http.Handle("/api/account", middleware.ValidateTokenMiddleware(http.HandlerFunc(controller_account.AccountController(db))))
}