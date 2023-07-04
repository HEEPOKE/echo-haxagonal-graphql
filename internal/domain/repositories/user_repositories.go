package repositories

import (
	"context"
	"fmt"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	DB *mongo.Database
}

func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	collection := r.DB.Collection("users")

	projection := bson.M{
		"password": 0,
	}

	cur, err := collection.Find(context.TODO(), bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var users []*models.User
	for cur.Next(context.TODO()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, cur.Err()
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

func (r *UserRepository) UpdateUser(user *models.User) error {
	collection := r.DB.Collection("users")

	filter := bson.M{"_id": user.ID}

	update := bson.M{
		"$set": bson.M{
			"username":  user.Username,
			"email":     user.Email,
			"tel":       user.Tel,
			"updatedAt": user.UpdatedAt,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := r.DB.Collection("users")

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
