package initialize

import (
	"context"
	"gluten/global"
	"gluten/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func MongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//在调用WithTimeout之后defer cancel()
	defer cancel()
	o := options.Client().ApplyURI("mongodb://" + global.MONGO.Host).SetAuth(
		options.Credential{
			AuthSource: "admin",
			Username:   global.MONGO.Username,
			Password:   global.MONGO.Password,
		})
	o.SetMaxPoolSize(20)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		log.Fatal(err.Error())
	}
	global.MongoDB = client.Database("gluten")
	service.InitGlutenInfoCollection()
	service.InitUserCategoryCollection()
	service.InitUserInfoCollection()
}
