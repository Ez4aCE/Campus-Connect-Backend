package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"campus-connect-backend/internal/models"
	"campus-connect-backend/internal/services"
)

type ClubHandler struct{
	service *services.ClubService
}

func NewClubHandler() *ClubHandler{
	return &ClubHandler{
		service:services.NewClubService() ,
	}
}

func GetUserID(c *gin.Context)(uuid.UUID, error){
	val, exists:=c.Get("userID")

	if !exists{
		return uuid.Nil, errors.New("userID not found in context")

	}

	userStr, ok:=val.(string)

	if !ok{
		return uuid.Nil, errors.New("userID is not a string")
	}

	return uuid.Parse(userStr)
}

func (h *ClubHandler) CreateClub(c *gin.Context){
	var req models.Club

	if err:=c.ShouldBindJSON(&req); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	userID, err:=GetUserID(c)

	if err != nil {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

	req.CreatedBy=userID

	if err:= h.service.CreateClub(&req); err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,req)
}

///clubs/:id/approve
func (h *ClubHandler) AprroveClub(c *gin.Context){
	clubID, err:=uuid.Parse(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid club id"})
		return
	}

	if err:= h.service.AprroveClub(clubID); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"club approved"})
}