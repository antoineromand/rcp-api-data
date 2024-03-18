package database

import (
	"rcp-api-data/internal/config/database/migration"
)


func InitDB() {
	db := InitPostgresDB()
	migration.RunMigration(db)
}