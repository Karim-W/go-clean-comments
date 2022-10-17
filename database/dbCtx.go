package database

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var instance *mongo.Client
var mtx = sync.Mutex{}

func GetDB() *mongo.Client {
	mtx.Lock()
	defer mtx.Unlock()
	if instance == nil {
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			panic(err)
		}
		err = client.Connect(context.Background())
		if err != nil {
			panic(err)
		}
		instance = client
		client.StartSession()
	}
	return instance
}
