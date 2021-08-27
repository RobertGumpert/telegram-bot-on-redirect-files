package config

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type (
	Mode string
	Type string
)

const (
	Production Mode = "production"
	Test       Mode = "test"
	Env        Type = "env"
	Json       Type = "json"
)

type Config struct {
	ApplicationHost string `json:"HOST"`
	ApplicationPort string `json:"PORT"`
	TelegramToken   string `json:"TGTOKEN"`
	DropboxToken    string `json:"DBTOKEN"`
}

func Reader(mode Mode, t Type) (Config, error) {
	var (
		path string
	)
	switch mode {
	case Test:
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return Config{}, err
		}
		path = dir
	case Production:
		_, dir, _, ok := runtime.Caller(0)
		if !ok {
			return Config{}, nil
		}
		dir = strings.Split(dir, "/src")[0]
		path = dir
	}
	switch t {
	case Env:
		return readEnv(path)
	}
	return Config{}, nil
}

func readEnv(path string) (Config, error) {
	var (
		env = new(Config)
	)
	err := os.Chdir(path)
	if err != nil {
		return *env, err
	}
	err = godotenv.Load()
	if err != nil {
		return *env, err
	}
	env.ApplicationPort = os.Getenv("PORT")
	env.ApplicationHost = os.Getenv("HOST")
	env.TelegramToken = os.Getenv("TGTOKEN")
	env.DropboxToken = os.Getenv("DBTOKEN")
	return *env, nil
}
