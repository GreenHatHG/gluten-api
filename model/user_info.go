package model

import (
	"github.com/jinzhu/gorm"
	"gluten/global"
)

type UserInfo struct {
	gorm.Model
	AvatarUrl string
	Username  string
	Email     string
	OauthId   int //oauth2登录后返回用户信息中的id字段
	Location  string
}

func (u UserInfo) CreateOrUpdateUserInfo() UserInfo {
	var query UserInfo
	global.DB.Where(UserInfo{OauthId: u.OauthId}).Assign(u).FirstOrCreate(&query)
	return query
}

//
//func GetUserByOpenId(openId string) (user *UserInfo, err error) {
//	user = new(UserInfo)
//	err = global.DB.Where(UserInfo{OpenId: openId}).First(user).Error
//	return
//}
