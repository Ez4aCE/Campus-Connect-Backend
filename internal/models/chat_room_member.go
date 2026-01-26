package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatRoomMember struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	RoomID uuid.UUID `gorm:"type:uuid;index;not null" json:"room_id"`
	UserID uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	JoinedAt time.Time `json:"joined_at"`
}