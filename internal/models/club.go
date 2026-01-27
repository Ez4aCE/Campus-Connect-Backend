package models

import (
	"time"

	"github.com/google/uuid"
)

type Club struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Status string `gorm:"type:varchar(20);not null;default:'pending';index" json:"status"`
	CreatedBy  uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
}