package model

import (
	"github.com/jinzhu/gorm"
	"gluten/global"
)

type UserInfo struct {
	gorm.Model
	AvatarUrl     string
	City          string
	Country       string
	NickName      string
	Province      string
	EncryptedData string
	Iv            string
	OpenId        string
	Code          string `gorm:"-"`
}

func AddUserInfo(info UserInfo) (err error) {
	err = global.DB.Create(&info).Error
	return
}

func GetUserByOpenId(openId string) (user *UserInfo, err error) {
	user = new(UserInfo)
	err = global.DB.Where(UserInfo{OpenId: openId}).First(user).Error
	return
}
