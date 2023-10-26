package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Email     string    `gorm:"not null;uniqueIndex"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null,default:now()"`
	UpdatedAt time.Time `gorm:"not null,default:now()"`
}
