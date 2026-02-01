package services


import (
	"errors"


	"campus-connect-backend/internal/db"
	"campus-connect-backend/internal/models"


	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrEventClosed = errors.New("event not open for registration")
	ErrAlreadyJoined = errors.New("already registred")
	ErrEventFull =  errors.New("event is full")
	ErrRegistrationNotFound = errors.New("registration not found")
)

type EventRegistrationService struct{
	db *gorm.DB
}

func NewEventRegistrationService() *EventRegistrationService{
	return &EventRegistrationService{
		db:db.DB,
	}
}


func (s *EventRegistrationService) Register(userID, eventID uuid.UUID) error{
	
	tx := s.db.Begin()

	var event models.Event

	if err:=tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&event,"id=?",eventID).Error;err!=nil{
			tx.Rollback()
			return ErrEventNotFound
		}
	if event.Status!="published"{
		tx.Rollback()
		return ErrEventClosed
	}

	if event.RegisteredCount>=event.Capacity{
		tx.Rollback()
		return ErrEventFull
	}

	var count int64

	if err:=tx.Model(&models.EventRegistration{}).
		Where("user_id=? AND event_id=?",userID, eventID).
		Count(&count).Error ; err!=nil{
			tx.Rollback()
			return err
		}

	if count>0{
		tx.Rollback()
		return ErrAlreadyJoined
	}

	reg:=models.EventRegistration{
		UserID: userID,
		EventID: eventID,
		Status: "registered",
	}

	if err:=tx.Create(&reg).Error; err!=nil{
		tx.Rollback()
		return err
	}
	event.RegisteredCount++

	if err:=tx.Save(&event).Error; err!=nil{
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *EventRegistrationService) Cancel(userID, eventID uuid.UUID) error{
	tx:=s.db.Begin()

	var reg models.EventRegistration

	if err:= tx.First(&reg,"user_id=? AND event_id=? AND status='registered'",
			userID,eventID).Error; err!=nil{
				tx.Rollback()
				return ErrRegistrationNotFound
			}

	if err:=tx.Delete(&reg).Error;err!=nil{
		tx.Rollback()
		return err
	}

	var event models.Event

	if err:=tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&event, "id=?",eventID).Error;err!=nil{
			tx.Rollback()
			return ErrEventNotFound
		}

	if event.RegisteredCount>0{
		event.RegisteredCount--
	}
	if err:=tx.Save(&event).Error;err!=nil{
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}