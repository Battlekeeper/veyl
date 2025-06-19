package types

import (
	"context"

	"github.com/Battlekeeper/veyl/internal/database"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Relay struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	PublicKey         string             `json:"public_key" bson:"public_key"`
	Name              string             `json:"name" bson:"name"`
	AuthenticationKey string             `json:"authentication_key" bson:"authentication_key"`

	// Associations
	Network primitive.ObjectID `json:"network" bson:"network"`
}

func GetRelayById(id primitive.ObjectID) (*Relay, error) {
	var relay Relay
	err := database.Client.Database("veyl").Collection("relays").FindOne(context.Background(), primitive.M{"_id": id}).Decode(&relay)
	if err != nil {
		return nil, err
	}
	return &relay, nil
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
