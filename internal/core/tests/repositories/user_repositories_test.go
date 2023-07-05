package repositories_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositorySuite struct {
	suite.Suite
	db *mongo.Database
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

func (suite *UserRepositorySuite) SetupSuite() {
	// Connect to the test MongoDB database
	db, err := setupDatabase()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to set up database: %s", err))
	}

	suite.db = db
}

func (suite *UserRepositorySuite) TearDownSuite() {
	// Clean up the test database
	suite.db.Drop(context.Background())
}

func setupDatabase() (*mongo.Database, error) {
	// Get the MongoDB connection string from an environment variable
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if connectionString == "" {
		return nil, fmt.Errorf("MONGODB_CONNECTION_STRING environment variable not set")
	}

	// Connect to the MongoDB database
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %s", err)
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %s", err)
	}

	// Get the test database from the MongoDB client
	db := client.Database("test_database")

	return db, nil
}

func TestGetAllUsers(t *testing.T) {
	// Set up the database connection
	db, err := setupDatabase()
	if err != nil {
		t.Fatalf("failed to set up database: %s", err)
	}

	// Create an instance of the UserRepository
	userRepo := &repositories.UserRepository{
		DB: db,
	}

	// Call the method being tested
	users, err := userRepo.GetAllUsers()

	// Assert the expected results
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestSaveUser(t *testing.T) {
	// Set up the database connection
	db, err := setupDatabase()
	if err != nil {
		t.Fatalf("failed to set up database: %s", err)
	}

	// Create an instance of the UserRepository
	userRepo := &repositories.UserRepository{
		DB: db,
	}

	// Create a mock user
	user := &models.User{
		ID:        primitive.NewObjectID().Hex(),
		Username:  "john_doe",
		Email:     "john@example.com",
		Tel:       "123456789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Call the method being tested
	err = userRepo.SaveUser(user)

	// Assert the expected results
	assert.NoError(t, err)
}

func TestGetUserByID(t *testing.T) {
	// Set up the database connection
	db, err := setupDatabase()
	if err != nil {
		t.Fatalf("failed to set up database: %s", err)
	}

	// Create an instance of the UserRepository
	userRepo := &repositories.UserRepository{
		DB: db,
	}

	// Create a mock user and save it to the database
	user := &models.User{
		ID:        primitive.NewObjectID().Hex(),
		Username:  "john_doe",
		Email:     "john@example.com",
		Tel:       "123456789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		t.Fatalf("failed to insert user: %s", err)
	}

	// Call the method being tested
	result, err := userRepo.GetUserByID(user.ID)

	// Assert the expected results
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)
}

func TestUpdateUser(t *testing.T) {
	// Set up the database connection
	db, err := setupDatabase()
	if err != nil {
		t.Fatalf("failed to set up database: %s", err)
	}

	// Create an instance of the UserRepository
	userRepo := &repositories.UserRepository{
		DB: db,
	}

	// Create a mock user and save it to the database
	user := &models.User{
		ID:        primitive.NewObjectID().Hex(),
		Username:  "john_doe",
		Email:     "john@example.com",
		Tel:       "123456789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		t.Fatalf("failed to insert user: %s", err)
	}

	// Update the user
	user.Username = "jane_doe"

	// Call the method being tested
	err = userRepo.UpdateUser(user)

	// Assert the expected results
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	// Set up the database connection
	db, err := setupDatabase()
	if err != nil {
		t.Fatalf("failed to set up database: %s", err)
	}

	// Create an instance of the UserRepository
	userRepo := &repositories.UserRepository{
		DB: db,
	}

	// Create a mock user and save it to the database
	user := &models.User{
		ID:        primitive.NewObjectID().Hex(),
		Username:  "john_doe",
		Email:     "john@example.com",
		Tel:       "123456789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		t.Fatalf("failed to insert user: %s", err)
	}

	// Call the method being tested
	err = userRepo.DeleteUser(user.ID)

	// Assert the expected results
	assert.NoError(t, err)
}
