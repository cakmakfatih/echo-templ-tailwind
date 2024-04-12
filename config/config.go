package config

import "github.com/joho/godotenv"

type AppConfig struct{}

func InitConfig() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}
