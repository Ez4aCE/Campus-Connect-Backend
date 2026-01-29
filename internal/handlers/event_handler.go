package handlers


import (
	"net/http"
	"time"


	"campus-connect-backend/internal/services"


	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"campus-connect-backend/internal/models"
)

type EventHandler struct{
	service *services.EventService
}

func NewEventHandler() *EventHandler{
	return &EventHandler{
		service: services.NewEventService(),
	}
}


type CreateEventRequest struct{
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	Location    string   `json:"location" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	Capacity    int `json:"capacity" binding:"required"`
}

//post /api/organizer/events
func (h *EventHandler) CreateGlobalEvent(c *gin.Context){
	var req CreateEventRequest

	if err:= c.ShouldBindJSON(&req); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	userID, err := GetUserID(c)

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}
	role:=c.MustGet("role").(string)

	event :=mapToEventModel(req)

	if err:= h.service.CreateGlobalHandler(userID, role, event); err!=nil{
		c.JSON(http.StatusForbidden,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,event)
}


//post /api/clubs/:id/events
func(h *EventHandler) CreateClubEvent(c *gin.Context){
	clubID, err:=uuid.Parse(c.Param("id"))

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid club id"})
		return
	}

	var req CreateEventRequest

	if err:=c.ShouldBindJSON(&req);  err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	userID, err:= GetUserID(c)
	role :=c.MustGet("role").(string)

	event := mapToEventModel(req)

	if err:= h.service.CreateClubEvent(userID,role,clubID,event) ; err!=nil{
		c.JSON(http.StatusForbidden, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusCreated, event)
}

//patch /api/events/:id/publish
func(h *EventHandler) PublishEvent(c *gin.Context){
	eventID, err:= uuid.Parse(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid event id"})
		return
	}
	userID, err:=GetUserID(c)
	role:=c.MustGet("role").(string)

	if err:= h.service.PublishEvent(userID, role, eventID); err!=nil{
		c.JSON(http.StatusForbidden, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"event published"})
}

//patch /api/events/:id/cancel
func (h *EventHandler) CancelEvent(c *gin.Context){
	eventID, err:= uuid.Parse(c.Param("id"))

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid event id"})
		return
	}
	userID, err:= GetUserID(c)
	role :=c.MustGet("role").(string)
	if err:= h.service.CancelEvent(userID,role,eventID) ; err!=nil{
		c.JSON(http.StatusForbidden,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"event cancelled"})

}


//get /api/events
func (h *EventHandler) ListEvents(c *gin.Context){
	events, err := h.service.ListEvents()
	if(err!=nil){
		c.JSON(http.StatusInternalServerError, gin.H{"error":"faild to fetch events list"})
		return
	}
	c.JSON(http.StatusOK,events)
}

//get /api/events/:id
func (h *EventHandler) GetEvent(c *gin.Context){
	eventID , err := uuid.Parse(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid event id"})
		return
	}
	event , err:=h.service.GetEvent(eventID)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"error":"event not found"})
		return
	}
	c.JSON(http.StatusOK,event)
}

func mapToEventModel(req CreateEventRequest) *models.Event{
	return &models.Event{
		Title: req.Title,
		Description: req.Description,
		Location: req.Location,
		StartTime: req.StartTime,
		EndTime:  req.EndTime,
		Capacity: req.Capacity,
	}
}