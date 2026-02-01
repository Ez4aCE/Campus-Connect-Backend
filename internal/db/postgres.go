package db

import (
	
	"log/slog"
	"campus-connect-backend/internal/config"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"campus-connect-backend/internal/models"
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

	err=DB.AutoMigrate(&models.User{},
					&models.Club{},
				&models.ClubMembership{},
			&models.Event{},
		&models.EventRegistration{},
	&models.Notification{},
&models.NotificationUser{})

	if err!=nil{
		slog.Error("Migration failed","error",err)
		panic(err)
	}
	slog.Info("Database migrated successfully")
}