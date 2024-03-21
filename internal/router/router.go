package router

import (
	"net/http"
	"rcp-api-data/internal/controller"
	"rcp-api-data/internal/middleware"

	"gorm.io/gorm"
)

func Router(db *gorm.DB) {
    http.Handle("/api/data", middleware.ValidateTokenMiddleware(http.HandlerFunc(controller.DataController())))
	// http.Handle("/api/data/insert", middleware.ValidateTokenMiddleware(http.HandlerFunc(controller.())))
}