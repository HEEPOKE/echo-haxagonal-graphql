package mocks

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDatabase(mongoURL string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %s", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %s", err)
	}

	db := client.Database("test_database")

	return db, nil
}

func TeardownDatabase(db *mongo.Database) error {
	err := db.Drop(context.Background())
	if err != nil {
		return err
	}

	return nil
}
