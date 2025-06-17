package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func StartClient() error {
	// Define context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Replace this URI with your MongoDB URI
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Ping the database
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	Client = client

	return nil
}
