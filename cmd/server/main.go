package main

import (
"log/slog"


"campus-connect-backend/internal/config"
"github.com/gin-gonic/gin"
"campus-connect-backend/internal/db"
)

func main(){
	config.Load()
	db.ConnectPostgres()
	r:=gin.Default()

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