package router

import (
	"net/http"
	"rcp-api-data/internal/controller"
)

func Router() {
	http.HandleFunc("/data", controller.DataController())
}