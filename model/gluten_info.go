package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type GlutenInfo struct {
	gorm.Model
	Title    string
	Star     int
	Post     datatypes.JSON
	Category string
	Company  datatypes.JSON
	UserId   uint
}

func (info GlutenInfo) String() string {
	data, _ := json.Marshal(info)
	return string(data)
}
