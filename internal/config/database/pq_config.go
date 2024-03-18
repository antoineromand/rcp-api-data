package database

import (
	"fmt"
	"rcp-api-data/internal/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitPostgresDB() *gorm.DB {
	env := config.GetConfig()
	if env == nil {
		return nil
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
    env.PG_HOST, env.PG_USER, env.PG_PASSWORD, env.PG_DBNAME, env.PG_PORT, "disable", "Europe/Paris")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxIdleTime(time.Hour)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}
	return db
}


func GetPostgresDB() *gorm.DB {
	return db
}