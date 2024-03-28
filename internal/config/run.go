package config

import (
	"rcp-api-data/internal/config/database"
	"rcp-api-data/internal/config/security"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RunConfig struct {
	Cfg *security.Environment
	Db  *gorm.DB
}

func Execute(sugar *zap.SugaredLogger, testing bool) *RunConfig {
	cfg, err := security.InitEnvironment(testing)
	if err != nil {
		sugar.Error("Config not found")
		return nil
	}
	sugar.Info("Config loaded")
	db, err := database.InitDB(cfg, sugar, cfg.MIGRATE)
	if err != nil {
		sugar.Error(err)
		return nil
	}
	sugar.Info("Database connected")
	return &RunConfig{
		Cfg: cfg,
		Db:  db,
	}
}
