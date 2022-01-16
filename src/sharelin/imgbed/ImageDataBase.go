package imgbed

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"main/sharelin/data/mongo"
	"main/sharelin/util"
)

const (
	database   = "imageBed"
	imageTable = "images"
)

var SnowflakeId, _ = util.CreateWorker(0, 0)

type ImageFile struct {
	Id         string `json:"id" bson:"_id"`
	Name       string `json:"name"`
	SourceName string `json:"sourceName"`
	Path       string `json:"path"`
	CreateTime string `json:"createTime"`
}

func initDatabase() {
	mongo.MONGO.Client.Database(database).CreateCollection(context.TODO(), imageTable)
}

var ImageDataBase = NewImageDataBaseOpt()

type imageDataBaseOpt struct {
}

func NewImageDataBaseOpt() *imageDataBaseOpt {
	initDatabase()
	return &imageDataBaseOpt{}
}

func (receiver imageDataBaseOpt) saveFileInfo(file ImageFile) {
	file.Id = SnowflakeId.NextId()
	collection := mongo.MONGO.Client.Database(database).Collection(imageTable)
	collection.InsertOne(context.TODO(), file)
}

func (receiver imageDataBaseOpt) queryFileList() *[]ImageFile {
	collection := mongo.MONGO.Client.Database(database).Collection(imageTable)
	filter := bson.D{}
	find, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("queryFileList acc err:%s\n", err.Error())
	}
	defer find.Close(context.TODO())
	var c []ImageFile
	find.All(context.TODO(), &c)
	return &c
}

func (receiver imageDataBaseOpt) queryOne(id string) (*ImageFile, error) {
	collection := mongo.MONGO.Client.Database(database).Collection(imageTable)

	filter := bson.M{"_id": id}
	one := collection.FindOne(context.TODO(), filter)
	if one.Err() != nil {
		log.Printf("queryOne acc err %s\n", one.Err().Error())
		return nil, one.Err()
	}
	var imageFile ImageFile
	err := one.Decode(&imageFile)
	if err != nil {
		log.Printf("queryOne acc err %s\n", one.Err().Error())
		return nil, err
	}
	return &imageFile, nil

}

func (receiver imageDataBaseOpt) deleteOne(id string) error {
	collection := mongo.MONGO.Client.Database(database).Collection(imageTable)
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("deleteOne acc err %s\n", err.Error())
		return err
	}
	return nil
}
