package routes


import (
"campus-connect-backend/internal/handlers"
"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine){
	api:= r.Group("/api")

	auth:=api.Group("/auth")
	{
		auth.POST("/register",handlers.Register)
		auth.POST("/login",handlers.Login)
	}
}