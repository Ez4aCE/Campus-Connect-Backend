package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Location    string    `json:"location"`
	Status      string    `gorm:"not null;default:'draft'" json:"status"`
	Budget      float64   `json:"budget"`
	CreatedBy  uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
}