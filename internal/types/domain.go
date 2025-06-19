package types

import (
	"context"

	"github.com/Battlekeeper/veyl/internal/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Domain struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`

	// Associations
	Networks []primitive.ObjectID `json:"networks" bson:"networks"`
	Owner    primitive.ObjectID   `json:"owner" bson:"owner"`
}

func CreateDomain(name string, owner primitive.ObjectID) Domain {
	domain := Domain{
		Id:       primitive.NewObjectID(),
		Name:     name,
		Networks: []primitive.ObjectID{},
		Owner:    owner,
	}
	// Assuming a database client is available
	database.Client.Database("veyl").Collection("domains").InsertOne(context.Background(), domain)
	return domain
}

func GetDomainById(id primitive.ObjectID) (*Domain, error) {
	var domain Domain
	err := database.Client.Database("veyl").Collection("domains").FindOne(
		context.Background(),
		primitive.M{"_id": id},
	).Decode(&domain)
	if err != nil {
		return nil, err
	}
	return &domain, nil
}

func GetDomainsByUserId(userId primitive.ObjectID) ([]Domain, error) {
	domains := make([]Domain, 0)
	cursor, err := database.Client.Database("veyl").Collection("domains").Find(
		context.Background(),
		primitive.M{"owner": userId},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var domain Domain
		if err := cursor.Decode(&domain); err != nil {
			return nil, err
		}
		domains = append(domains, domain)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return domains, nil
}

func (u *Domain) Update() error {
	_, err := database.Client.Database("veyl").Collection("domains").UpdateOne(
		context.Background(),
		primitive.M{"_id": u.Id},
		primitive.M{"$set": u},
	)
	return err
}

func (u *Domain) AddNetwork(id primitive.ObjectID) error {
	u.Networks = append(u.Networks, id)
	return u.Update()
}

func (u *Domain) RemoveNetwork(id primitive.ObjectID) error {
	for i, networkId := range u.Networks {
		if networkId == id {
			u.Networks = append(u.Networks[:i], u.Networks[i+1:]...)
			return u.Update()
		}
	}
	return nil
}

func (u *Domain) GetNetworks() ([]veylNetwork, error) {
	var networks = make([]veylNetwork, 0)
	if len(u.Networks) == 0 {
		return networks, nil
	}
	cursor, err := database.Client.Database("veyl").Collection("networks").Find(
		context.Background(),
		primitive.M{"_id": primitive.M{"$in": u.Networks}},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var network veylNetwork
		if err := cursor.Decode(&network); err != nil {
			return nil, err
		}
		networks = append(networks, network)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return networks, nil
}

func DeleteDomain(id primitive.ObjectID) error {
	domain, err := GetDomainById(id)
	if err != nil {
		return err
	}
	if domain == nil {
		return nil
	}

	// Remove all networks associated with the domain
	for _, networkId := range domain.Networks {
		err = DeleteNetwork(networkId)
		if err != nil {
			return err
		}
	}

	// Delete the domain itself
	_, err = database.Client.Database("veyl").Collection("domains").DeleteOne(context.Background(), primitive.M{"_id": id})
	return err
}
