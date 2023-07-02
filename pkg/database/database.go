package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoDB(uri, databaseName string) (*MongoDB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")

	db := client.Database(databaseName)
	return &MongoDB{
		client:   client,
		database: db,
	}, nil
}

func (db *MongoDB) Close() error {
	err := db.client.Disconnect(context.Background())
	if err != nil {
		return err
	}

	log.Println("Disconnected from MongoDB")
	return nil
}

func (db *MongoDB) GetDatabase() *mongo.Database {
	return db.database
}
