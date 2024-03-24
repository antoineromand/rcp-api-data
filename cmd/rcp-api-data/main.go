package main

import (
	"flag"
	"net/http"
	"rcp-api-data/internal/config"
	"rcp-api-data/internal/router"

	"rcp-api-data/internal/utils"
)

func main() {
	migration := flag.Bool("migration", false, "Run migration")
	flag.Parse()
	sugar := utils.GetLogger()
	sugar.Info("Starting server...")
	sugar.Info("Database connected")
	var runConfig *config.RunConfig
	runConfig = config.Execute(sugar, migration, false)
	db := runConfig.Db
	cfg := runConfig.Cfg
	if runConfig == nil {
		sugar.Error("Could not run configurations")
		return
	}
	router.Router(db, cfg)
	err := http.ListenAndServe(":" + cfg.Port, nil)
	if err != nil {
		sugar.Error(err)
		return
	}
}