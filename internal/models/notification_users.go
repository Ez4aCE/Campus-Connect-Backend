package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationUser struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	NotificationID uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:uniq_notification_user" json:"notification_id"`
	UserID         uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:uniq_notification_user" json:"user_id"`

	IsRead bool `gorm:"not null;default:false;index" json:"is_read"`

	ReadAt *time.Time `json:"read_at,omitempty"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime;index" json:"created_at"`

	Notification Notification `gorm:"foreignKey:NotificationID"`
}


func (NotificationUser) TableName() string {
	return "notification_users"
}