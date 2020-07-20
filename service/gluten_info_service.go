package service

import (
	"context"
	. "github.com/gobeam/mongo-go-pagination"
	"gluten/global"
	"gluten/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/datatypes"
)

var glutenInfoCollection *mongo.Collection

func InitGlutenInfoCollection() {
	if glutenInfoCollection == nil {
		glutenInfoCollection = global.MongoDB.Collection("gluten_info")
	}
}

func AddGlutenInfo(info model.GlutenInfo) error {
	_, err := glutenInfoCollection.InsertOne(context.TODO(), info)
	return err
}

func SelectAllGlutenInfoById(userId string, currentPage int64, pageSize int64, sort string) (error, []model.GlutenInfo) {
	paginatedData, err := New(glutenInfoCollection).Limit(pageSize).Page(currentPage).Sort(sort, -1).Filter(bson.D{
		{"user_id", userId},
	}).Find()
	if err != nil {
		return err, nil
	}
	var lists []model.GlutenInfo
	for _, raw := range paginatedData.Data {
		var gluten *model.GlutenInfo
		marshallErr := bson.Unmarshal(raw, &gluten)
		var hexId struct {
			ID primitive.ObjectID `bson:"_id"`
		}
		idErr := bson.Unmarshal(raw, &hexId)
		if marshallErr == nil && idErr == nil {
			gluten.ID = hexId.ID.Hex()
			lists = append(lists, *gluten)
		}
	}
	return nil, lists
}

func CountGlutenInfo(userId string) (int64, error) {
	count, err := glutenInfoCollection.CountDocuments(context.TODO(), bson.D{
		{"user_id", userId},
	})
	return count, err
}

func UpdateGlutenInfoTitle(id string, title string, userId string) error {
	filter := getCommonFilter(id, userId)
	update := bson.D{
		{"$set", bson.D{
			{"title", title},
		}},
	}

	_, err := glutenInfoCollection.UpdateOne(context.TODO(), filter, update)
	return err
}

func UpdateGlutenInfoValue(id string, userId string, key string, value datatypes.JSON) error {
	filter := getCommonFilter(id, userId)

	update := bson.D{
		{"$set", bson.D{
			{key, value},
		}},
	}
	_, err := glutenInfoCollection.UpdateOne(context.TODO(), filter, update)
	return err
}

func UpdateGlutenInfoStar(id string, userId string, star int) error {
	filter := getCommonFilter(id, userId)

	update := bson.D{
		{"$set", bson.D{
			{"star", star},
		}},
	}
	_, err := glutenInfoCollection.UpdateOne(context.TODO(), filter, update)
	return err
}

func DeleteGlutenById(id string, userId string) error {
	filter := getCommonFilter(id, userId)
	_, err := glutenInfoCollection.DeleteOne(context.TODO(), filter)
	return err
}

func SelectGlutenById(id string, userId string) (result model.GlutenInfo, err error) {
	hexId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{
		{"_id", hexId},
		{"user_id", userId},
	}
	err = glutenInfoCollection.FindOne(context.TODO(), filter).Decode(&result)
	return
}

func getCommonFilter(id string, userId string) bson.D {
	hexId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{
		{"_id", hexId},
		{"user_id", userId},
	}
	return filter
}
