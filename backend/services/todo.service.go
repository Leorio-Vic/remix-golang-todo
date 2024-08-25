package services

import (
	"context"
	"fmt"
	"log"
	"rgt-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoService struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewTodoService(client *mongo.Client) *TodoService {
	db := client.Database("rgt")

	collectionName := "todos"
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

	collection := db.Collection(collectionName)
	return &TodoService{
		db:         db,
		collection: collection,
	}
}

func (s *TodoService) CreateTodo(ctx context.Context, todo *models.Todos) (*mongo.InsertOneResult, error) {
	return s.collection.InsertOne(ctx, todo)
}

func (s *TodoService) GetTodo(ctx context.Context) ([]*models.Todos, error) {
	cursor, err := s.collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}
	var todo []*models.Todos

	err = cursor.All(ctx, &todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
