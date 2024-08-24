package db

import (
	"context"
    "fmt"
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
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	createCollection(client)
}

func createCollection (client *mongo.Client) {
	collectionName := "todos"

	db = client.Database("rgt")
	collections, err := db.ListCollectionNames(context.TODO(), map[string]interface{}{})

	if err != nil {
		log.Fatal(err)
	}

	collectionExists := false
    for _, name := range collections {
        if name == collectionName {
            collectionExists = true
            break
        }
    }

	if !collectionExists {
        // Tạo collection nếu nó chưa tồn tại
        err = db.CreateCollection(context.TODO(), collectionName)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Collection created:", collectionName)
    } else {
        fmt.Println("Collection already exists:", collectionName)
    }
}