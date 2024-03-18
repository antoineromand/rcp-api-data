package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)
type Environment struct {
	Port string `env:"PORT"`
	RCP_AUTH_URL string `env:"RCP_AUTH_URL"`
	RCP_AUTH_PORT string `env:"RCP_AUTH_PORT"`
	RCP_AUTH_PROTOCOL string `env:"RCP_AUTH_PROTOCOL"`
	PG_USER string `env:"DB_SERVER_USER"`
	PG_PASSWORD string `env:"DB_SERVER_PASSWORD"`
	PG_DBNAME string `env:"DB_SERVER_DATABASE"`
	PG_PORT string `env:"DB_SERVER_PORT"`
	PG_HOST string `env:"DB_SERVER_HOST"`
}

var cfg Environment

func init() {
	loadEnv := godotenv.Load()
	if loadEnv != nil {
		fmt.Print(loadEnv)
	}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Erreur lors du parsing des variables d'environnement : %+v\n", err)
	}
}

func (e *Environment) GetAuthURL() string {
	return e.RCP_AUTH_PROTOCOL + "://" + e.RCP_AUTH_URL + ":" + e.RCP_AUTH_PORT
}

func GetConfig() *Environment {
	return &cfg
}

