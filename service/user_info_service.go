package service

import (
	"context"
	"gluten/global"
	"gluten/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var userInfoCollection *mongo.Collection

func InitUserInfoCollection() {
	if userInfoCollection == nil {
		userInfoCollection = global.MongoDB.Collection("user_info")
	}
}

func CreateOrUpdateUserInfo(u model.UserInfo) (model.UserInfo, error) {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"oauth_id", u.OauthId}}
	update := bson.D{
		{"$setOnInsert", bson.D{{"create_at", time.Now().Local()}}},
		{"$set", bson.D{
			{"updated_at", time.Now().Local()},
			{"avatar_url", u.AvatarUrl},
			{"email", u.Email},
			{"location", u.Location},
			{"username", u.Username},
		}}}
	if _, err := userInfoCollection.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		return model.UserInfo{}, err
	} else {
		err = userInfoCollection.FindOne(context.TODO(), filter).Decode(&u)
		return u, err
	}
}
