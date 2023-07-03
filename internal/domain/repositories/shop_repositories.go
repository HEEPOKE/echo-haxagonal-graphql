package repositories

import (
	"context"
	"time"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShopRepository struct {
	DB *mongo.Database
}

func (r *ShopRepository) GetShopByID(id string) (*models.Shop, error) {
	collection := r.DB.Collection("shops")

	var shop models.Shop

	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&shop)
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
