package config

import (
	"os"
	"log/slog"
	"github.com/joho/godotenv"
)

type Config struct{
	AppName string
	AppPort string
	DatabaseURL string
} 

var AppConfig Config

func Load(){
	err:=godotenv.Load();

	if err!=nil{
		slog.Info("No .env file found, using system env")
	}

	AppConfig=Config{
		AppName: os.Getenv("APP_NAME"),
		AppPort: os.Getenv("APP_PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}