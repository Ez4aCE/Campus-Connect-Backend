package routes

import (
	"github.com/gin-gonic/gin"

	"campus-connect-backend/internal/handlers"
	"campus-connect-backend/internal/middleware"
)

func RegisterClubRoutes(r *gin.RouterGroup){
	clubHandler:= handlers.NewClubHandler()
	memberHandler:= handlers.NewMembershipHandler();

	clubs:= r.Group("/clubs")

	clubs.Use(middleware.JWTAuth())


	//ORGANIZER CREATES CLUB
	clubs.POST("/",middleware.RequireRoles("organizer"),clubHandler.CreateClub)

	//admin approve club
	clubs.PATCH("/:id/approve",middleware.RequireRoles("admin"),clubHandler.AprroveClub)

	//join club
	clubs.POST("/:id/join",memberHandler.RequestJoin)

	//approve join
	clubs.PATCH("/:id/members/:uid/approve",memberHandler.ApproveMember)


	//leave member
	clubs.DELETE("/:id/leave",memberHandler.LeaveClub)
}