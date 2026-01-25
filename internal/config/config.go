package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Config struct{
	AppName string
	AppPort string
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
	DBSSLMODE string
} 

var AppConfig Config

func Load(){
	err:=godotenv.Load();

	if err!=nil{
		log.Println("No .env file found, using system env")
	}

	AppConfig=Config{
		AppName: os.Getenv("APP_NAME"),
		AppPort: os.Getenv("APP_PORT"),
		DBHost:  os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBSSLMODE: os.Getenv("DB_SSLMODE"),
	}
}