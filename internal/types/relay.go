package types

import (
	"context"
	"fmt"

	"github.com/Battlekeeper/veyl/internal/database"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Relay struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	PublicKey string             `json:"public_key" bson:"public_key"`
	Name      string             `json:"name" bson:"name"`

	// Associations
	NetworkId primitive.ObjectID `json:"network_id" bson:"network_id"`
}

func GetRelayById(id primitive.ObjectID) (*Relay, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$unwind", Value: "$relays"}},
		{{Key: "$match", Value: bson.D{{Key: "relays._id", Value: id}}}},
		{{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: "$relays"}}}},
	}

	cursor, err := database.Client.Database("veyl").Collection("networks").Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate relays: %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var relay Relay
		if err := cursor.Decode(&relay); err != nil {
			return nil, fmt.Errorf("failed to decode relay: %v", err)
		}
		if relay.Id == id {
			return &relay, nil
		}
	}

	return nil, fmt.Errorf("relay with id %s not found", id.Hex())
}

type RelayAuth struct {
	RelayID   string `json:"relayid"`
	PublicKey string `json:"public_key"`
	IP        string `json:"ip"`
	Port      int    `json:"port"`
}

type RelayClient struct {
	Auth       RelayAuth       `json:"auth"`
	Connection *websocket.Conn `json:"-"`
}

type RelayConnection struct {
	RelayID   string `json:"relayid"`
	PublicKey string `json:"public_key"`
	IP        string `json:"ip"`
	Port      int    `json:"port"`
}
