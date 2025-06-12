package types

import (
	"context"

	"github.com/Battlekeeper/veyl/internal/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           primitive.ObjectID   `json:"id" bson:"_id"`
	Domains      []primitive.ObjectID `json:"domains" bson:"domains"`
	Email        string               `json:"email" bson:"email"`
	PasswordHash string               `json:"password_hash" bson:"password_hash"`
}

func GetUserById(id primitive.ObjectID) (*User, error) {
	var user User
	err := database.Client.Database("veyl").Collection("users").FindOne(context.Background(), primitive.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Update() error {
	_, err := database.Client.Database("veyl").Collection("users").UpdateOne(
		context.Background(),
		primitive.M{"_id": u.Id},
		primitive.M{"$set": u},
	)
	return err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(email, passwordRaw string) (*User, error) {
	hashed, err := HashPassword(passwordRaw)
	if err != nil {
		return nil, err
	}
	user := &User{
		Id:           primitive.NewObjectID(),
		Domains:      []primitive.ObjectID{},
		Email:        email,
		PasswordHash: hashed,
	}

	_, err = database.Client.Database("veyl").Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
