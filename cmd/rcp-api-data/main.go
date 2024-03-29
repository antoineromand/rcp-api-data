package main

import (
	"net/http"
	"rcp-api-data/internal/config"
	"rcp-api-data/internal/router"

	"rcp-api-data/internal/utils"
)

func main() {
	sugar := utils.GetLogger()
	sugar.Info("Starting server...")
	sugar.Info("Database connected")
	runConfig := config.Execute(sugar, false)
	db := runConfig.Db
	if db == nil {
		sugar.Error("Could not connect to database")
		return
	}
	cfg := runConfig.Cfg
	if cfg == nil {
		sugar.Error("Could not load configurations")
		return
	}
	if runConfig == nil {
		sugar.Error("Could not run configurations")
		return
	}
	router.Router(db, cfg, sugar)
	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		sugar.Error(err)
		return
	}
}
