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

var userCategoryCollection *mongo.Collection

func InitUserCategoryCollection() {
	if userCategoryCollection == nil {
		userCategoryCollection = global.MongoDB.Collection("user_category")
	}
}

//func CreateOrUpdateUserCategory(u model.UserCategory) error {
//	result, err := SelectUserCategoryById(u.UserId)
//	//没有记录则创建，否则更新
//	if err != nil{
//		u.CreatedAt = time.Now().Local()
//		_, err := userCategoryCollection.InsertOne(context.TODO(), u)
//		if err != nil {
//			return err
//		}
//	}else{
//		//判断内容不一样才更新
//		if !testEq(result.Company, u.Company) || !testEq(result.Post, u.Post) || !testEq(result.Category, u.Category){
//			u.CreatedAt = time.Now().Local()
//			update := bson.D{
//				{"$set", bson.D{
//					{"category", u.Category},
//					{"post", u.Post},
//					{"company", u.Company},
//				}},
//			}
//			_, err := userCategoryCollection.UpdateOne(context.TODO(),  bson.D{{"userid", u.UserId}}, update)
//			if err != nil{
//				return err
//			}
//		}
//	}
//	return nil
//}

func CreateOrUpdateUserCategory(u model.UserCategory) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"user_id", u.UserId}}
	update := bson.D{
		{"$setOnInsert", bson.D{{"create_at", time.Now().Local()}}},
		{"$set", bson.D{
			{"updated_at", time.Now().Local()},
			{"category", u.Category},
			{"post", u.Post},
			{"company", u.Company},
		}}}
	_, err := userCategoryCollection.UpdateOne(context.TODO(), filter, update, opts)
	return err
}

func SelectUserCategoryById(id string) (result model.UserCategory, err error) {
	err = userCategoryCollection.FindOne(context.TODO(), bson.D{{"user_id", id}}).Decode(&result)
	return
}

//func testEq(a, b [] string) bool {
//
//	// If one is nil, the other must also be nil.
//	if (a == nil) != (b == nil) {
//		return false;
//	}
//	if len(a) != len(b) {
//		return false
//	}
//	for i := range a {
//		if a[i] != b[i] {
//			return false
//		}
//	}
//	return true
//}
