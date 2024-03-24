package database

import (
	"errors"
	"rcp-api-data/internal/config/database/migration"
	"rcp-api-data/internal/config/security"

	"go.uber.org/zap"
	"gorm.io/gorm"
)


func InitDB(env *security.Environment, sugar *zap.SugaredLogger, _migration *bool) (*gorm.DB, error) {
	db := InitPostgresDB(env)
	if db == nil {
		return nil, errors.New("database initialization failed")
	}
	if *_migration {
		migration.RunMigration(db, sugar)
	} else {
		sugar.Info("Migration skipped")
	}
	return db, nil
}