package lib

import "go.mongodb.org/mongo-driver/mongo"

var DbClient *mongo.Client

func SetDBClient(client *mongo.Client) {
	DbClient = client
}

func Collections(collectionName string)  *mongo.Collection {
	collection := Db().Collection(collectionName)
	return collection
}

func Db() *mongo.Database {
	db := DbClient.Database("metaverse")
	return db
}