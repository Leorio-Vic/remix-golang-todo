package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todos struct {
	ID        primitive.ObjectID
	Title     string
	Completed bool
}
