package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"main/config"
	"time"
)

type MongoDatabase struct {
	client *mongo.Client
}

func (mongoDatabase *MongoDatabase) connectMongo() {
	clientOptions := options.Client().ApplyURI(config.Config("mongodb.address"))
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		client, err = mongo.Connect(context.TODO(), clientOptions)
	}
	mongoDatabase.client = client

}

func (mongoDatabase *MongoDatabase) StartCheckConnection() {
	go func() {
		for true {
			if mongoDatabase.client == nil {
				mongoDatabase.connectMongo()
			}
			// 检查连接
			err := mongoDatabase.client.Ping(context.TODO(), nil)
			if err != nil {
				log.Println(err.Error())
				if mongoDatabase.client != nil {
					mongoDatabase.client.Disconnect(context.TODO())
				}
				mongoDatabase.connectMongo()
			}
			time.Sleep(5)
		}
	}()
}
