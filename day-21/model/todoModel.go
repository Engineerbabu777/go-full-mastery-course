package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty"`
	Status string             `json:"status,omitempty"`
}
