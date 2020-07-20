package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserInfo struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	UpdatedAt time.Time          `bson:"update_at"`
	DeletedAt time.Time          `bson:"deleted_at"`
	AvatarUrl string             `bson:"avatar_url"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	OauthId   int                `bson:"oauth_id"` //oauth2登录后返回用户信息中的id字段
	Location  string             `bson:"location"`
}

//
//func GetUserByOpenId(openId string) (user *UserInfo, err error) {
//	user = new(UserInfo)
//	err = global.DB.Where(UserInfo{OpenId: openId}).First(user).Error
//	return
//}
