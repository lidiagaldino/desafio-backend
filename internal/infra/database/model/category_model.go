package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string  `bson:"title"`
	Description string  `bson:"description"`
	OwnerID     string  `bson:"owner_id"`
}
