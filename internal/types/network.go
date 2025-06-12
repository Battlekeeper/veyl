package types

import (
	"context"

	"github.com/Battlekeeper/veyl/internal/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type veylNetwork struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`

	// Associations
	Relays    []Relay            `bson:"relays"`
	Resources []Relay            `bson:"resources"`
	Domain    primitive.ObjectID `json:"domain" bson:"domain"`
}

func CreateNetwork(name string) veylNetwork {
	network := veylNetwork{
		Id:        primitive.NewObjectID(),
		Name:      name,
		Relays:    []Relay{},
		Resources: []Relay{},
	}
	database.Client.Database("veyl").Collection("networks").InsertOne(context.Background(), network)
	return network
}

func GetNetworkById(id primitive.ObjectID) (*veylNetwork, error) {
	var network veylNetwork
	err := database.Client.Database("veyl").Collection("networks").FindOne(context.Background(), primitive.M{"_id": id}).Decode(&network)
	if err != nil {
		return nil, err
	}
	return &network, nil
}

func (vn *veylNetwork) Update() error {
	_, err := database.Client.Database("veyl").Collection("networks").UpdateOne(
		context.Background(),
		primitive.M{"_id": vn.Id},
		primitive.M{"$set": vn},
	)
	return err
}

func (vn *veylNetwork) AddRelay(relay Relay) error {
	vn.Relays = append(vn.Relays, relay)
	return vn.Update()
}

func (vn *veylNetwork) AddResource(resource Relay) error {
	vn.Resources = append(vn.Resources, resource)
	return vn.Update()
}

func (vn *veylNetwork) GetDomain() (Domain, error) {
	var domain Domain
	err := database.Client.Database("veyl").Collection("domains").FindOne(context.Background(), primitive.M{"_id": vn.Domain}).Decode(&domain)
	if err != nil {
		return Domain{}, err
	}
	return domain, nil
}
