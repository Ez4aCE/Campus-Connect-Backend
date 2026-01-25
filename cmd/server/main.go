package main

import (
"log"


"campus-connect-backend/internal/config"
"github.com/gin-gonic/gin"
)

func main(){
	config.Load()

	r:=gin.Default()

	r.GET("/health", func (c *gin.Context){ 
		c.JSON(200,gin.H{
			"status":"ok",
			"app":config.AppConfig.AppName,
			"port":config.AppConfig.AppPort,
		})
	})
	
	log.Println("Starting server on port ",config.AppConfig.AppPort)
	r.Run(":"+config.AppConfig.AppPort)

}