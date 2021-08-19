package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host    string
	TgToken string
	AppPort string
}

func ReadENV() (Config, error) {
	var (
		env = new(Config)
	)
	err := godotenv.Load()
	if err != nil {
		return *env, err
	}
	env.AppPort = os.Getenv("PORT")
	env.Host = os.Getenv("HOST")
	env.TgToken = os.Getenv("TGTOKEN")
	return *env, nil
}
