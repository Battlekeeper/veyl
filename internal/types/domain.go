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
}

func CreateDomain(name string) Domain {
	domain := Domain{
		Id:       primitive.NewObjectID(),
		Name:     name,
		Networks: []primitive.ObjectID{},
	}
	// Assuming a database client is available
	database.Client.Database("veyl").Collection("domains").InsertOne(context.Background(), domain)
	return domain
}

func (u *Domain) Update() error {
	_, err := database.Client.Database("veyl").Collection("domains").UpdateOne(
		context.Background(),
		primitive.M{"_id": u.Id},
		primitive.M{"$set": u},
	)
	return err
}
