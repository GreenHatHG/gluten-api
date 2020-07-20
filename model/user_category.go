package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserCategory struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	CreatedAt time.Time          `bson:"create_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Category  []string
	Company   []string
	Post      []string
	UserId    string `bson:"user_id"`
}
