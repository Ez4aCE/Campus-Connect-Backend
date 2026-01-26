package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID   uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	Title    string    `gorm:"not null" json:"title"`
	Message  string    `json:"message"`
	IsRead   bool      `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}