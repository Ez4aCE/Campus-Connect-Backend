package models

import "github.com/google/uuid"

type Resource struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Type        string    `gorm:"not null" json:"type"`
	Description string    `json:"description"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
}