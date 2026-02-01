package routes


import (
	"campus-connect-backend/internal/handlers"
	"campus-connect-backend/internal/middleware"


	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(r *gin.RouterGroup){
	
	handler := handlers.NewEventHandler()
	regHandler:=handlers.NewEventRegistrationHandler()

	events:=r.Group("/events")
	events.Use(middleware.JWTAuth())
	{
		events.GET("/",handler.ListEvents)
		events.GET("/:id", handler.GetEvent)

		events.PATCH("/:id/publish",handler.PublishEvent)
		events.PATCH("/:id/cancel", handler.CancelEvent)

		events.POST("/:id/register",regHandler.RegisterEvent)
		events.DELETE("/:id/register",regHandler.CancelRegister)

	}

	organizer:=r.Group("/organizer")
	organizer.Use(middleware.JWTAuth(),middleware.RequireRoles("organizer","admin"))
	{
		organizer.POST("/events",handler.CreateGlobalEvent)
	}
	clubs:=r.Group("/clubs")
	clubs.Use(middleware.JWTAuth())
	{
		clubs.POST("/:id/events",handler.CreateClubEvent)
	}
}