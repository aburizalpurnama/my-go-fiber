package domain

import (
	"time"
)

type User struct {
	Id        string    `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Email     string    `gorm:"not null;uniqueIndex"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
}
