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

func TestGetAllShops(t *testing.T) {
	mockRepo := new(mocks.ShopRepositoryMock)
	expectedShops := []*models.Shop{
		{
			ID:        primitive.NewObjectID().Hex(),
			Name:      "Shop 1",
			Address:   "123 Main St",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        primitive.NewObjectID().Hex(),
			Name:      "Shop 2",
			Address:   "456 Elm St",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	mockRepo.On("GetAllShops", mock.Anything).Return(expectedShops, nil)

	shops, err := mockRepo.GetAllShops(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedShops, shops)

	mockRepo.AssertCalled(t, "GetAllShops", mock.Anything)
}

func TestGetShopByID(t *testing.T) {
	mockRepo := new(mocks.ShopRepositoryMock)
	expectedShop := &models.Shop{
		ID:        primitive.NewObjectID().Hex(),
		Name:      "Shop 1",
		Address:   "123 Main St",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockRepo.On("GetShopByID", mock.Anything, expectedShop.ID).Return(expectedShop, nil)

	shop, err := mockRepo.GetShopByID(context.Background(), expectedShop.ID)

	assert.NoError(t, err)
	assert.Equal(t, expectedShop, shop)

	mockRepo.AssertCalled(t, "GetShopByID", mock.Anything, expectedShop.ID)
}

func TestCreateShop(t *testing.T) {
	mockRepo := new(mocks.ShopRepositoryMock)
	shop := &models.Shop{
		Name:      "Shop 1",
		Address:   "123 Main St",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockRepo.On("CreateShop", mock.Anything, shop).Return(nil)

	err := mockRepo.CreateShop(context.Background(), shop)

	assert.NoError(t, err)

	mockRepo.AssertCalled(t, "CreateShop", mock.Anything, shop)
}

func TestUpdateShop(t *testing.T) {
	mockRepo := new(mocks.ShopRepositoryMock)
	shop := &models.Shop{
		ID:        "1",
		Name:      "Updated Shop",
		Address:   "456 Main St",
		UpdatedAt: time.Now(),
	}
	mockRepo.On("UpdateShop", mock.Anything, shop).Return(nil)

	err := mockRepo.UpdateShop(context.Background(), shop)

	assert.NoError(t, err)

	mockRepo.AssertCalled(t, "UpdateShop", mock.Anything, shop)
}

func TestDeleteShop(t *testing.T) {
	mockRepo := new(mocks.ShopRepositoryMock)
	shopID := "1"
	mockRepo.On("DeleteShop", mock.Anything, shopID).Return(nil)

	err := mockRepo.DeleteShop(context.Background(), shopID)

	assert.NoError(t, err)

	mockRepo.AssertCalled(t, "DeleteShop", mock.Anything, shopID)
}
