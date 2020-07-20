package service

import (
	. "github.com/gobeam/mongo-go-pagination"
	"gluten/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SelectGlutenInfoByCategory(userId string, currentPage int64, pageSize int64, sort string, category string) (error, []model.GlutenInfo) {
	paginatedData, err := New(glutenInfoCollection).Limit(pageSize).Page(currentPage).Sort(sort, -1).Filter(bson.D{
		{"user_id", userId},
		{"category", category},
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
