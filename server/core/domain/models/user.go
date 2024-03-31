package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	IsAdmin   bool      `json:"is_admin" gorm:"default:false;not null"`
	IsActive  bool      `json:"is_active" gorm:"default:true;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp;not null"`
}
