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
	NetworkId primitive.ObjectID `json:"network_id" bson:"network_id"`
}

func GetResourceById(id primitive.ObjectID) (*Resource, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$unwind", Value: "$resources"}},
		{{Key: "$match", Value: bson.D{{Key: "resources._id", Value: id}}}},
		{{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: "$resources"}}}},
	}

	cursor, err := database.Client.Database("veyl").Collection("networks").Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate resources: %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var resource Resource
		if err := cursor.Decode(&resource); err != nil {
			return nil, fmt.Errorf("failed to decode resources: %v", err)
		}
		if resource.Id == id {
			return &resource, nil
		}
	}

	return nil, fmt.Errorf("resources with id %s not found", id.Hex())
}
