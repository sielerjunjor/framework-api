package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Framework struct {
	ID          *primitive.ObjectID `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Type  string        `bson:"type" json:"type"`
	Description string        `bson:"description" json:"description"`
}
