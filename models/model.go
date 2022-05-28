package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	Body     string             `json:"body"`
	Author   string             `json:"author"`
	Comments []Comment          `json:"comments,omitempty"`
}

type Comment struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}
