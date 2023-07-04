package test

import (
	"context"
	"testing"
	"time"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/mocks"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllUsers(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	expectedUsers := []*models.User{
		{
			ID:        "1",
			Username:  "user1",
			Email:     "user1@example.com",
			Tel:       "123456789",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "2",
			Username:  "user2",
			Email:     "user2@example.com",
			Tel:       "987654321",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	mockRepo.On("GetAllUsers", mock.Anything).Return(expectedUsers, nil)

	users, err := mockRepo.GetAllUsers(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)

	mockRepo.AssertCalled(t, "GetAllUsers", mock.Anything)
}

func TestGetUserByID(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	expectedUser := &models.User{
		ID:        primitive.NewObjectID().Hex(),
		Username:  "user1",
		Email:     "user1@example.com",
		Password:  nil,
		Tel:       "123456789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockRepo.On("GetUserByID", mock.Anything, expectedUser.ID).Return(expectedUser, nil)

	user, err := mockRepo.GetUserByID(context.Background(), expectedUser.ID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	mockRepo.AssertCalled(t, "GetUserByID", mock.Anything, expectedUser.ID)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	user := &models.User{
		Username:  "user1",
		Email:     "user1@example.com",
		Password:  nil,
		Tel:       "123456789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockRepo.On("CreateUser", mock.Anything, user).Return(nil)

	err := mockRepo.CreateUser(context.Background(), user)

	assert.NoError(t, err)

	mockRepo.AssertCalled(t, "CreateUser", mock.Anything, user)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	user := &models.User{
		ID:        primitive.NewObjectID().Hex(),
		Username:  "user1",
		Email:     "user1@example.com",
		Password:  nil,
		Tel:       "123456789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockRepo.On("UpdateUser", mock.Anything, user).Return(nil)

	err := mockRepo.UpdateUser(context.Background(), user)

	assert.NoError(t, err)

	mockRepo.AssertCalled(t, "UpdateUser", mock.Anything, user)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	userID := primitive.NewObjectID().Hex()
	mockRepo.On("DeleteUser", mock.Anything, userID).Return(nil)

	err := mockRepo.DeleteUser(context.Background(), userID)

	assert.NoError(t, err)

	mockRepo.AssertCalled(t, "DeleteUser", mock.Anything, userID)
}
