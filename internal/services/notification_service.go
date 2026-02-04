package services

import (
	"campus-connect-backend/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationService struct{
	db *gorm.DB
}

func NewNotificationService(db *gorm.DB) *NotificationService{
	return &NotificationService{
		db: db,
	}
}

func (s *NotificationService) CreateForUsers(
	title string,
	message string,
	notificationType string,
	metadata []byte,
	createdBy *uuid.UUID,
	userIDs []uuid.UUID,
)error{
	tx :=s.db.Begin()
	defer tx.Rollback()

	notification :=models.Notification{
		ID: uuid.New(),
		Type: notificationType,
		Title: title,
		Message: message,
		Metadata: metadata,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(30*24*time.Hour),
	}

	if err:=tx.Create(&notification).Error; err!=nil{
		return err
	}

	var notifUsers []models.NotificationUser

	for _,uid:=range userIDs{
		notifUsers = append(notifUsers, models.NotificationUser{
			ID: uuid.New(),
			NotificationID: notification.ID,
			UserID:  uid,
			IsRead : false,
			CreatedAt:  time.Now(),
		})
	}

	if err:=tx.Create(&notifUsers).Error; err!=nil{
		return err
	}
	return tx.Commit().Error

}

func (s *NotificationService) CreateForUser(
	title, message, notificationType string,
	metadata []byte,
	createdBy *uuid.UUID,
	userID uuid.UUID,
) error{
	return s.CreateForUsers(
		title,
		message,
		notificationType,
		metadata,
		createdBy,
		[]uuid.UUID{userID},
	)
}

func (s *NotificationService) Broadcast(
	title , message, notificationType string,
	metadata []byte,
	createdBy *uuid.UUID,
) error{
	var users []uuid.UUID
	if err := s.db.Model(&models.User{}).Pluck("id", &users).Error ;err!=nil{
		return err
	}

	return s.CreateForUsers(title,message,notificationType,
		metadata,
		createdBy,users)
}


func (s *NotificationService) MarkRead(userID , notificationID uuid.UUID) error{
	tx:=s.db.Begin()
	defer tx.Rollback()

	now := time.Now()

	if err:=tx.Model(&models.NotificationUser{}).
		Where("user_id =? AND notification_id=?", userID, notificationID).
		Updates(map[string]interface{}{
			"is_read":true,
			"read_at":now,
		}).Error; err!=nil{
			return err
		}

	return tx.Commit().Error
}

func (s *NotificationService) MarkAllRead(userID uuid.UUID) error{
	tx:= s.db.Begin()
	defer tx.Rollback()

	now:= time.Now()

	if err:= tx.Model(&models.NotificationUser{}).
		Where("user_id=? AND is_read=false",userID).
		Updates(map[string]interface{}{
			"is_read":true,
			"read_at":now,
		}).Error ; err!=nil{
			return err
		}

	return tx.Commit().Error
}

func(s *NotificationService) ListByUser(userID uuid.UUID, limit, offset int)(
	[]models.NotificationUser, error,
){
	var results []models.NotificationUser

	err:=s.db.Preload("Notification").
		Where("user_id=?",userID).Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&results).Error

	return results, err
}


func (s *NotificationService) UnreadCount(userID uuid.UUID)(int64, error){
	var count int64

	err:=s.db.Model(&models.NotificationUser{}).
		Where("user_id=? AND is_read=false",userID).
		Count(&count).Error

	return count, err
}