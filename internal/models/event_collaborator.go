package models

import "github.com/google/uuid"

type EventCollaborator struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	EventID uuid.UUID `gorm:"type:uuid;index;not null" json:"event_id"`
	ClubID  uuid.UUID `gorm:"type:uuid;index;not null" json:"club_id"`
	Role    string    `gorm:"not null" json:"role"`
}