package db

import (
	"context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TodosCollection *mongo.Collection
var db *mongo.Database
func InitDB() {
	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
		NilMapAsEmpty:     true,
		NilSliceAsEmpty:   true,
	}

	clientOpts := options.Client().
		ApplyURI("mongodb://localhost:27017").
		SetBSONOptions(bsonOpts)

	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	db = client.Database("rgt")

	listCollectionNames, err := db.ListCollectionNames(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collectionExists := false
	for _, name := range listCollectionNames {
		if name == "todos" {
			collectionExists = true
			break
		}
	}

	if !collectionExists {
		err = db.CreateCollection(ctx, "todos")
		if err != nil {
			log.Fatal(err)
		}
	}

	TodosCollection = db.Collection("todos")
}
