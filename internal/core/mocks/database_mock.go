package mocks

import (
	"context"
	"fmt"

	"github.com/HEEPOKE/echo-haxagonal-graphql/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDatabase() (*mongo.Database, error) {
	connectionString := config.Cfg.MONGO_URL
	if connectionString == "" {
		return nil, fmt.Errorf("MONGODB_CONNECTION_STRING environment variable not set")
	}

	clientOptions := options.Client().ApplyURI(connectionString)
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
