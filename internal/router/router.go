package router

import (
	"net/http"
	"rcp-api-data/internal/controller"
	"rcp-api-data/internal/middleware"
)

func Router() {
    http.Handle("/data", middleware.ValidateTokenMiddleware(http.HandlerFunc(controller.DataController())))
}