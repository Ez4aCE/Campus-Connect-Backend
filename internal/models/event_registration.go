package models

import (
	"time"

	"github.com/google/uuid"
)

type EventRegistration struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	EventID uuid.UUID `gorm:"type:uuid;index;not null" json:"event_id"`
	UserID  uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	Status  string    `gorm:"not null;default:'registered'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
}