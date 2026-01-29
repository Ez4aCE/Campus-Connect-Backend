package services

import (
	"errors"

	"campus-connect-backend/internal/db"
	"campus-connect-backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrUnauthorizedEvent = errors.New("unauthorized event action")
	ErrInvalidEventState = errors.New("Invalid event state")
	ErrInvalidTimeRange = errors.New("invalid time range")
	ErrInvalidCapacity = errors.New("invalid capacity")
	ErrEventNotFound = errors.New("event not found")
)

type EventService struct{
	db *gorm.DB
}

func NewEventService() *EventService{
	return &EventService{
		db: db.DB,
	}
}

func (s *EventService) CreateGlobalHandler(userID uuid.UUID, role string, event *models.Event) error{
	if role!="admin" && role!="organizer"{
		return ErrUnauthorizedEvent
	}

	if err:=validateEvent(event); err!=nil{
		return err
	}

	event.ID=uuid.Nil
	event.ClubID=nil
	event.CreatedBy=userID
	event.Status="draft"
	event.RegisteredCount=0

	return s.db.Create(event).Error
}

func (s *EventService) CreateClubEvent(userID uuid.UUID, role string, clubID uuid.UUID, event *models.Event) error{
	if role!="admin"{
		var membership models.ClubMembership

		err := s.db.Where(
				"user_id = ? AND club_id = ? AND role = 'admin' AND status = 'approved'",
				userID, clubID,
				).First(&membership).Error


		if err != nil {
			return ErrUnauthorizedEvent
		}
	}

	if err:= validateEvent(event);err!=nil{
		return err
	}

	event.ID=uuid.Nil
	event.ClubID=&clubID
	event.CreatedBy=userID
	event.Status="draft"
	event.RegisteredCount=0;

	return s.db.Create(event).Error
}


func (s *EventService) PublishEvent(userID uuid.UUID, role string, eventID uuid.UUID) error{

	var event models.Event
	if err:= s.db.First(&event,"id=?",eventID).Error; err!=nil{
		return ErrEventNotFound
	}

	if event.Status!="draft"{
		return ErrInvalidEventState
	}

	if !s.canManageEvent(userID,role,&event){
		return ErrUnauthorizedEvent
	}

	event.Status="published"
	return s.db.Save(&event).Error
}

func (s *EventService) CancelEvent(userID uuid.UUID, role string, eventID uuid.UUID) error{
	var event models.Event

	if err:= s.db.First(&event,"id=?",eventID).Error ; err!=nil{
		return ErrEventNotFound
	}

	if event.Status=="completed"{
		return ErrInvalidEventState
	}

	if !s.canManageEvent(userID,role,&event){
		return ErrUnauthorizedEvent
	}

	event.Status="cancelled"
	return s.db.Save(&event).Error
}


func (s *EventService) ListEvents()([]models.Event,error){
	var events []models.Event

	err := s.db.Order("start_time asc").Find(&events).Error
	return events, err
}

func (s *EventService) GetEvent(eventID uuid.UUID)(*models.Event, error){
	var event models.Event

	err:=s.db.First(&event,"id=?",eventID).Error
	return &event, err
}

func (s *EventService) canManageEvent(userID uuid.UUID, role string, event *models.Event) bool{
	if role=="admin"{
		return true;
	}
	if event.ClubID==nil{
		return event.CreatedBy==userID
	}

	var membership models.ClubMembership

	err := s.db.Where("user_id = ? AND club_id = ? AND role = 'admin' AND status = 'approved'",
				userID,*event.ClubID,).First(&membership).Error

	return err==nil
}

func validateEvent(event *models.Event) error{
	if event.Title == "" || event.Location == ""{
		return errors.New("missing required fields")
	}

	if event.EndTime.Before(event.StartTime){
		return ErrInvalidTimeRange
	}

	if event.Capacity<=0{
		return ErrInvalidCapacity
	}

	return nil
}