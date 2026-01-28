package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatRoom struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Type   string    `gorm:"not null" json:"type"` // private, club, event
	RefID uuid.UUID `gorm:"type:uuid;index" json:"ref_id"`
	CreatedAt time.Time `json:"created_at"`
}