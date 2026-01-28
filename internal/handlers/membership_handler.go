package handlers


import (
	"net/http"


	"github.com/gin-gonic/gin"
	"github.com/google/uuid"


	"campus-connect-backend/internal/services"
)

type MembershipHandler struct{
	service *services.MembershipService
}

func NewMembershipHandler() *MembershipHandler{
	return &MembershipHandler{
		service: services.NewMembershipService(),
	}
}


///clubs/:id/join
func (h *MembershipHandler) RequestJoin(c *gin.Context){
	userID, err:=GetUserID(c)

	if err != nil {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }
	

	clubID, err:= uuid.Parse(c.Param("id"))

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid club id"})
		return
	}

	if err:=h.service.RequestJoin(userID,clubID); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"join request sent"})
}

// clubs/:id/members/:uid/approve
func (h *MembershipHandler) ApproveMember( c *gin.Context){
	adminID, err:=GetUserID(c)

	if err != nil {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

	clubID, err:= uuid.Parse(c.Param("id"))

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid club id"})
		return
	}

	userID, err := uuid.Parse(c.Param("uid"))

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid user id"})
		return
	}

	if err:= h.service.ApproveMember(adminID,clubID,userID); err!=nil{
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"member approved"})
}

//clubs/:id/leave
func (h *MembershipHandler) LeaveClub(c *gin.Context){

	userID, err:=GetUserID(c)

	if err != nil {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

	clubID, err:= uuid.Parse(c.Param("id"))

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid club id"})
		return
	}
	
	if err:=h.service.LeaveClub(userID,clubID);err!=nil{
		c.JSON(http.StatusConflict, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"left club"})
}