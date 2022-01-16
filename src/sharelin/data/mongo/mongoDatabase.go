package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"main/config"
	"time"
)

type Database struct {
	uri    string
	Client *mongo.Client
}

func (mongoDatabase *Database) connectMongo() {
	clientOptions := options.Client().ApplyURI(mongoDatabase.uri)
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		client, err = mongo.Connect(context.TODO(), clientOptions)
	}
	mongoDatabase.Client = client
}

func InitDataBase(uri string) *Database {
	dataBase := Database{uri: uri}
	dataBase.connectMongo()
	dataBase.startCheckConnection()
	return &dataBase
}

func (mongoDatabase *Database) startCheckConnection() {
	go func() {
		for true {
			if mongoDatabase.Client == nil {
				mongoDatabase.connectMongo()
			}
			// 检查连接
			err := mongoDatabase.Client.Ping(context.TODO(), nil)
			if err != nil {
				log.Println(err.Error())
				if mongoDatabase.Client != nil {
					mongoDatabase.Client.Disconnect(context.TODO())
				}
				mongoDatabase.connectMongo()
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

const mongoUri = "mongodb.uri"

var MONGO = InitDataBase(config.CONFIG.Get(mongoUri))
