package main

import (
	"log/slog"
	"time"

	"campus-connect-backend/internal/config"
	"campus-connect-backend/internal/db"
	"campus-connect-backend/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	config.Load()
	db.ConnectPostgres()
	r:=gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{"GET","POST","PUT","PATCH","DELETE","OPTIONS"},
		AllowHeaders: []string{"Origin","Content-Type","Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12* time.Hour,
	}))

	routes.RegisterRoutes(r)
	r.GET("/health", func (c *gin.Context){ 
		c.JSON(200,gin.H{
			"status":"ok",
			"app":config.AppConfig.AppName,
			"port":config.AppConfig.AppPort,
		})
	})
	
	slog.Info("Starting server","port",config.AppConfig.AppPort)
	r.Run(":"+config.AppConfig.AppPort)

}