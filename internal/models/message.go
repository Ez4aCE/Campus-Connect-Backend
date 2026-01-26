package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	RoomID  uuid.UUID `gorm:"type:uuid;index;not null" json:"room_id"`
	SenderID uuid.UUID `gorm:"type:uuid;index;not null" json:"sender_id"`
	Content string    `gorm:"not null" json:"content"`
	CreatedAt time.Time `json:"created_at"`
}