package types

import (
	"context"
	"fmt"
	"net"

	"github.com/Battlekeeper/veyl/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Resource struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Alias   string             `json:"alias" bson:"alias"`
	Address net.IP             `json:"address" bson:"address"`

	// Associations
	Network primitive.ObjectID `json:"network" bson:"network"`
}

func GetResourceById(id primitive.ObjectID) (*Resource, error) {
	var resource Resource
	err := database.Client.Database("veyl").Collection("resources").FindOne(context.Background(), bson.M{"_id": id}).Decode(&resource)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("resource not found")
		}
		return nil, fmt.Errorf("error retrieving resource: %v", err)
	}
	return &resource, nil
}
