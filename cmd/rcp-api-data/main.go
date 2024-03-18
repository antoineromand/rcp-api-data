package main

import (
	"fmt"
	"net/http"

	"rcp-api-data/internal/config"
	"rcp-api-data/internal/config/database"
	"rcp-api-data/internal/router"
)

func main() {
	
	fmt.Println("Starting server...")
	database.InitDB()
	fmt.Println("Database initialized...")
	cfg := config.GetConfig()
	if cfg == nil {
		fmt.Println("Erreur lors de la récupération de la configuration")
		return
	}
	router.Router()
	err := http.ListenAndServe(":" + cfg.Port, nil)
	if err != nil {
		fmt.Println(err)
	}
}