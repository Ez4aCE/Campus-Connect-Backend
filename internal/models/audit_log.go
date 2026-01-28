package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID   uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	Action   string    `gorm:"not null" json:"action"`
	Entity   string    `gorm:"not null" json:"entity"`
	EntityID uuid.UUID `gorm:"type:uuid;index" json:"entity_id"`
	CreatedAt time.Time `json:"created_at"`
}