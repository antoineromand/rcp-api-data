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

