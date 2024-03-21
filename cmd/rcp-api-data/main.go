package main

import (
	"flag"
	"net/http"
	"rcp-api-data/internal/config"
	"rcp-api-data/internal/config/database"
	"rcp-api-data/internal/router"

	"rcp-api-data/internal/utils"
)

func main() {
	migration := flag.Bool("migration", false, "Run migration")
	flag.Parse()
	sugar := utils.GetLogger()
	sugar.Info("Starting server...")
	db, errDB := database.InitDB(migration)
	if errDB != nil {
		sugar.Error(errDB)
		return
	}
	sugar.Info("Database connected")
	cfg := config.GetConfig()
	if cfg == nil {
		sugar.Error("Config not found")
		return
	}
	router.Router(db)
	err := http.ListenAndServe(":" + cfg.Port, nil)
	if err != nil {
		sugar.Error(err)
		return
	}
}