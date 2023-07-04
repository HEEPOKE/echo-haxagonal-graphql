package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
	"github.com/stretchr/testify/assert"
)

func TestShopService_GetShopByID(t *testing.T) {
	// Mock dependencies
	shopRepo := &mockShopRepository{}
	shopService := NewShopService(shopRepo)

	// Define test data
	shopID := "shop-123"
	expectedShop := &models.Shop{
		ID:        shopID,
		Name:      "Test Shop",
		Address:   "This is a test shop",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Mock the behavior of the ShopRepository
	shopRepo.On("GetShopByID", shopID).Return(expectedShop, nil)

	// Call the method being tested
	result, err := shopService.GetShopByID(context.Background(), shopID)

	// Assert the expected result
	assert.NoError(t, err)
	assert.Equal(t, expectedShop, result)

	// Assert that the mocked method was called
	shopRepo.AssertCalled(t, "GetShopByID", shopID)
}

func TestShopService_CreateShop(t *testing.T) {
	// Mock dependencies
	shopRepo := &mockShopRepository{}
	shopService := NewShopService(shopRepo)

	// Define test data
	newShop := &models.Shop{
		ID:        "new-shop-123",
		Name:      "New Shop",
		Address:   "This is a new shop",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Mock the behavior of the ShopRepository
	shopRepo.On("CreateShop", newShop).Return(nil)

	// Call the method being tested
	err := shopService.CreateShop(context.Background(), newShop)

	// Assert the expected result
	assert.NoError(t, err)

	// Assert that the mocked method was called
	shopRepo.AssertCalled(t, "CreateShop", newShop)
}
