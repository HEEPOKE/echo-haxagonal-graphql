package repositories_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/mocks"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/repositories"
	"github.com/HEEPOKE/echo-haxagonal-graphql/pkg/config"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositorySuite struct {
	suite.Suite
	db *mongo.Database
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

func (suite *UserRepositorySuite) SetupSuite() {
	mongoURL := config.Cfg.MONGO_URL
	db, err := mocks.SetupDatabase(mongoURL)
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to set up mongo database: %s", err))
	}

	suite.db = db
}

func (suite *UserRepositorySuite) TearDownSuite() {
	mocks.TeardownDatabase(suite.db)
}

func (suite *UserRepositorySuite) TestGetAllUsers() {
	userRepo := &repositories.UserRepository{
		DB: suite.db,
	}

	users, err := userRepo.GetAllUsers()

	suite.NoError(err)
	suite.NotNil(users)
}

func (suite *UserRepositorySuite) TestSaveUser() {
	userRepo := &repositories.UserRepository{
		DB: suite.db,
	}

	user := &models.User{
		ID:        primitive.NewObjectID().Hex(),
		Username:  "john_doe",
		Email:     "johnq@example.com",
		Tel:       "123456789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := userRepo.SaveUser(user)

	suite.NoError(err)
}

func (suite *UserRepositorySuite) TestGetUserByID() {
	userRepo := &repositories.UserRepository{
		DB: suite.db,
	}

	user := &models.User{
		ID:        primitive.NewObjectID().Hex(),
		Username:  "john_doe",
		Email:     "johnaaaa@example.com",
		Tel:       "123456789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := suite.db.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to insert user: %s", err))
	}

	result, err := userRepo.GetUserByID(user.ID)

	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(user.ID, result.ID)
}

func (suite *UserRepositorySuite) TestUpdateUser() {
	userRepo := &repositories.UserRepository{
		DB: suite.db,
	}

	user := &models.User{
		ID:        primitive.NewObjectID().Hex(),
		Username:  "john_doe",
		Email:     "john@example.com",
		Tel:       "123456789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := suite.db.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to insert user: %s", err))
	}

	user.Username = "jane_doe"

	err = userRepo.UpdateUser(user)

	suite.NoError(err)
}
