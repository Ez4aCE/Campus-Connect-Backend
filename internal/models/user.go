package models

import(
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type User struct{
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string `gorm:"not null"`
	Email string `gorm:"not null"`
	Password string `gorm:"not null"`
	Role string  `gorm:"not null;default:'participant'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error{
	u.ID=uuid.New()
	return nil
}