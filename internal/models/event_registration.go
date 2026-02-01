package models

import (
	"time"

	"github.com/google/uuid"
)

type EventRegistration struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

	UserID  uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_event" json:"user_id"`
	EventID uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_event" json:"event_id"`

	Status string `gorm:"type:varchar(20);not null;default:'registered';index" json:"status"`

	CreatedAt time.Time `json:"created_at"`
}