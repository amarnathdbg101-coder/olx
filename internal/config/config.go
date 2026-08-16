package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
	DatabaseUrl string
}

func MustLoad() Config {
	_ = godotenv.Load()
	port := os.Getenv("PORT")
	if port == ""{
		panic("port is required")
	}

	env := os.Getenv("ENV")
	if env == ""{
		panic("ENV is required")
	}
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == ""{
		panic("Database url is required")
	}
	return Config{
		Port: port,
		Env: env,
		DatabaseUrl: databaseUrl,
	}
}