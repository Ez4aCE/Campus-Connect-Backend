package models

import (
	"time"

	"github.com/google/uuid"
)

type ResourceBooking struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ResourceID uuid.UUID  `gorm:"type:uuid;index;not null" json:"resource_id"`
	UserID    uuid.UUID  `gorm:"type:uuid;index;not null" json:"user_id"`
	EventID   *uuid.UUID `gorm:"type:uuid" json:"event_id,omitempty"`
	StartTime time.Time  `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`
	Status    string     `gorm:"not null;default:'pending'" json:"status"`
}