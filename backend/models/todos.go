package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todos struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title     string `bson:"title,omitempty" json:"title,omitempty"`
	Completed bool `bson:"completed,omitempty" json:"completed,omitempty"`
}
