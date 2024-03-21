package database

import (
	"errors"
	"rcp-api-data/internal/config/database/migration"
	"rcp-api-data/internal/utils"

	"gorm.io/gorm"
)


func InitDB(_migration *bool) (*gorm.DB, error) {
	sugar := utils.GetLogger()
	db := InitPostgresDB()
	if db == nil {
		return nil, errors.New("database initialization failed")
	}
	if *_migration {
		migration.RunMigration(db)
	} else {
		sugar.Info("Migration skipped")
	}
	return db, nil
}