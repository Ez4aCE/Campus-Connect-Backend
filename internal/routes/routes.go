package routes


import (
"campus-connect-backend/internal/handlers"
"github.com/gin-gonic/gin"
	"campus-connect-backend/internal/middleware"
)

func RegisterRoutes(r *gin.Engine){
	api:= r.Group("/api")

	auth:=api.Group("/auth")
	{
		auth.POST("/register",handlers.Register)
		auth.POST("/login",handlers.Login)
	}

	protected:=api.Group("/")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/me",handlers.Me)
		admin:=protected.Group("/admin")
		admin.Use(middleware.RequireRoles("admin"))
		{
			admin.GET("/dashboard",handlers.AdminDashboard)
		}
		organizer := protected.Group("/organizer") 
		organizer.Use(middleware.RequireRoles("organizer", "admin"))
		{ 
			organizer.POST("/events", handlers.CreateEvent) 
		}
	}
}