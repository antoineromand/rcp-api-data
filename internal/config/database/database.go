package database

import (
	"rcp-api-data/internal/config/database/migration"
)


func InitDB(_migration *bool) {
	db := InitPostgresDB()
	if *_migration {
		migration.RunMigration(db)
	} else {
		println("Migration skipped")
	}
}