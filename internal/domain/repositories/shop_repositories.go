package repositories

import (
	"context"
	"time"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShopRepository struct {
	DB *mongo.Database
}

func (r *ShopRepository) GetAllShops() ([]*models.Shop, error) {
	collection := r.DB.Collection("shops")

	projection := bson.M{
		"password": 0,
	}

	cur, err := collection.Find(context.TODO(), bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var shops []*models.Shop
	for cur.Next(context.TODO()) {
		var shop models.Shop
		err := cur.Decode(&shop)
		if err != nil {
			return nil, err
		}
		shops = append(shops, &shop)
	}

	return shops, cur.Err()
}

func (r *ShopRepository) GetShopByID(id string) (*models.Shop, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := r.DB.Collection("shops")

	filter := bson.M{"_id": objectID}

	var shop models.Shop

	err = collection.FindOne(context.TODO(), filter).Decode(&shop)
	if err != nil {
		return nil, err
	}

	return &shop, nil
}

func (r *ShopRepository) CreateShop(shop *models.Shop) error {
	collection := r.DB.Collection("shops")
	shop.CreatedAt = time.Now()
	shop.UpdatedAt = time.Now()

	_, err := collection.InsertOne(context.TODO(), shop)
	if err != nil {
		return err
	}

	return nil
}
