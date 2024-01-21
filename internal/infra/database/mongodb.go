package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
 
	if err != nil {
	 log.Fatal(err)
	}
 
	err = client.Ping(context.Background(), nil)
	if err != nil {
	 log.Fatal(err)
	}
 
	fmt.Println("Successfully connected to MongoDB")
	return client, nil
}
