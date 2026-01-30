package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Location    string `gorm:"not null" json:"location"`

	StartTime time.Time `gorm:"not null" json:"start_time"`
	EndTime   time.Time `gorm:"not null" json:"end_time"`

	Capacity int `gorm:"not null;check:capacity > 0" json:"capacity"`
	RegisteredCount int `gorm:"not null;default:0;check:registered_count >= 0" json:"registered_count"`

	Status string `gorm:"type:varchar(30);not null;default:'draft';index" json:"status"`

	ClubID   *uuid.UUID `gorm:"type:uuid;index" json:"club_id,omitempty"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null;index" json:"created_by"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}