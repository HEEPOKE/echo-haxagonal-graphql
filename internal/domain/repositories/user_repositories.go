package repositories

import (
	"context"
	"fmt"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := r.DB.Collection("users")

	filter := bson.M{"_id": objectID}

	var user models.User
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}
