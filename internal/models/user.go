package models

import(
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type User struct{
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Name          string     `gorm:"not null" json:"name"`
	Email         string     `gorm:"unique;not null" json:"email"`
	PasswordHash  string     `gorm:"not null" json:"-"`
	Role          string     `gorm:"not null;default:'participant'" json:"role"`
	
	Department    string     `json:"department"`
	Year          int        `json:"year"`
	Phone         string     `json:"phone"`
	IsVerfied	  bool       `gorm:"default:false" json:"is_verified"`
	
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error{
	u.ID=uuid.New()
	return nil
}