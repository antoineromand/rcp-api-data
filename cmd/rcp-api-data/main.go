package main

import (
	"fmt"
	"net/http"
	Config "rcp-api-data/internal/config"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")
	loadEnv := godotenv.Load()
	if loadEnv != nil {
		fmt.Print(loadEnv)
	}
	cfg := Config.Environment{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println("Port:", cfg.Port)
	err := http.ListenAndServe(":" + cfg.Port, nil)
	if err != nil {
		fmt.Println(err)
	}
}