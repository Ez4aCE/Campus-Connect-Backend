package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Notification struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Type      string         `gorm:"size:50;not null;index" json:"type"`
	Title     string         `gorm:"not null" json:"title"`
	Message   string         `gorm:"not null" json:"message"`
	Metadata  datatypes.JSON `gorm:"type:jsonb" json:"metadata,omitempty"`

	CreatedBy *uuid.UUID `gorm:"type:uuid" json:"created_by,omitempty"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"created_at"`
	ExpiresAt time.Time `gorm:"index" json:"expires_at"`

	Users []NotificationUser `gorm:"foreignKey:NotificationID" json:"-"`
}

func (Notification) TableName() string {
	return "notifications"
}