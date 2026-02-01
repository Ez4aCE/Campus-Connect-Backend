package handlers


import (
	"net/http"


	"campus-connect-backend/internal/services"


	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type EventRegistrationHandler struct{
	service *services.EventRegistrationService
}

func NewEventRegistrationHandler() *EventRegistrationHandler{
	return &EventRegistrationHandler{
		service: services.NewEventRegistrationService(),
	}
}


//post /api/events/:id/register
func (h *EventRegistrationHandler) RegisterEvent(c *gin.Context){
	eventID, err:=uuid.Parse(c.Param("id"))

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid event id"})
		return
	}
	userID, err:=GetUserID(c)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if err:= h.service.Register(userID,eventID); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return

	}
	c.JSON(http.StatusCreated, gin.H{"message":"registred succesfully"})

}

//delete /api/events/:id/register

func (h *EventRegistrationHandler) CancelRegister(c *gin.Context){
	eventID, err := uuid.Parse(c.Param("id"))

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid event id"})
		return
	}

	userID, err:= GetUserID(c)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	if err:=h.service.Cancel(userID,eventID); err!=nil{
		
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	
	}

	c.JSON(http.StatusOK, gin.H{"message":"registration cancelled"})
}