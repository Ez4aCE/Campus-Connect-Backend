package db

import (
	
	"log/slog"
	"campus-connect-backend/internal/config"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectPostgres(){
	

	dsn:= config.AppConfig.DatabaseURL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!=nil{
		slog.Error("Failed to connect to PostgreSQL","ERROR",err)
		panic(err)
	}
	DB=db
	slog.Info("Postgres connected succesfully")
}