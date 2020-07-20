package model

import (
	"encoding/json"
	"gorm.io/datatypes"
	"time"
)

type GlutenInfo struct {
	ID        string    `bson:"-"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	DeletedAt time.Time `bson:"deleted_at"`
	Title     string
	Star      int
	Post      datatypes.JSON
	Category  []string
	Company   datatypes.JSON
	UserId    string `bson:"user_id"`
}

func (info GlutenInfo) String() string {
	data, _ := json.Marshal(info)
	return string(data)
}
