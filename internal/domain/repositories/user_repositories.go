package repositories

import (
	"context"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
)

type UserRepository struct {
	DB *mongo.Database
}

func (r *UserRepository) SaveUser(user *models.User) error {
	collection := r.DB.Collection("users")

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	collection := r.DB.Collection("users")

	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
