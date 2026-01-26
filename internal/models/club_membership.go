package models

import (
	"time"

	"github.com/google/uuid"
)

type ClubMembership struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID   uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	ClubID   uuid.UUID `gorm:"type:uuid;index;not null" json:"club_id"`
	Role     string    `gorm:"not null;default:'member'" json:"role"`
	JoinedAt time.Time `json:"joined_at"`
}