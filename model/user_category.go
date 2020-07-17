package model

import (
	"time"
)

type UserCategory struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Category  string
	Company   string
	Post      string
}
