package model

import (
	"gluten/global"
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

func (u UserCategory) CreateOrUpdateUserCategory() UserCategory {
	var query UserCategory
	global.DB.Where(UserCategory{ID: u.ID}).Assign(u).FirstOrCreate(&query)
	return query
}

func SelectUserCategoryById(id uint) UserCategory {
	var query UserCategory
	global.DB.First(&query, UserCategory{ID: id})
	return query
}
