package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

func MustLoad() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("godotenv:%v",err)
	}
	port := os.Getenv("PORT")
	if port == ""{
		panic("port is required")
	}

	env := os.Getenv("ENV")
	if env == ""{
		panic("ENV is required")
	}
}