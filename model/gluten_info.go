package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"gluten/global"
	"gorm.io/datatypes"
)

type GlutenInfo struct {
	gorm.Model
	Title    string
	Star     int
	Post     datatypes.JSON
	Category int
	Company  datatypes.JSON
	UserId   int
}

func (info GlutenInfo) String() string {
	data, _ := json.Marshal(info)
	return string(data)
}

func AddGlutenInfo(info GlutenInfo) {
	if err := global.DB.Create(&info).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("插入成功")
	}
}

func SelectAllGlutenInfo() (err error, info []GlutenInfo) {
	err = global.DB.Find(&info).Error
	return
}
