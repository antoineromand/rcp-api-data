package security

import (
	"errors"
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Environment struct {
	Port              string `env:"PORT"`
	RCP_AUTH_URL      string `env:"RCP_AUTH_URL"`
	RCP_AUTH_PORT     string `env:"RCP_AUTH_PORT"`
	RCP_AUTH_PROTOCOL string `env:"RCP_AUTH_PROTOCOL"`
	RCP_AUTH_PREFIX   string `env:"RCP_AUTH_PREFIX"`
	PG_USER           string `env:"DB_SERVER_USER"`
	PG_PASSWORD       string `env:"DB_SERVER_PASSWORD"`
	PG_DBNAME         string `env:"DB_SERVER_DATABASE"`
	PG_PORT           string `env:"DB_SERVER_PORT"`
	PG_HOST           string `env:"DB_SERVER_HOST"`
	CORS_ORIGIN       string `env:"CORS_ORIGIN"`
	MIGRATE           bool   `env:"MIGRATE"`
	PREFIX            string `env:"PREFIX"`
	RP_BROKERS        string `env:"RP_BROKERS"`
}

func InitEnvironment(testing bool) (*Environment, error) {
	var loadEnv error
	if testing {
		loadEnv = godotenv.Load("../.env")
	} else {
		loadEnv = godotenv.Load()
	}
	if loadEnv != nil {
		fmt.Print(loadEnv)
	}
	cfg := Environment{}
	if err := env.Parse(&cfg); err != nil {
		err = errors.New("error while parsing env variables")
		return nil, err
	}
	return &cfg, nil
}

func (e *Environment) GetAuthURL() string {
	url := e.RCP_AUTH_PROTOCOL + "://" + e.RCP_AUTH_URL + ":" + e.RCP_AUTH_PORT
	if len(e.RCP_AUTH_PREFIX) > 0 {
		url += "/" + e.RCP_AUTH_PREFIX
	}
	return url
}
