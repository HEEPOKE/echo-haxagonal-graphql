package mocks

import (
	"context"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type ShopRepositoryMock struct {
	mock.Mock
	DB DatabaseInterface
}

func (m *ShopRepositoryMock) GetAllShops(ctx context.Context) ([]*models.Shop, error) {
	args := m.Called()
	return args.Get(0).([]*models.Shop), args.Error(1)
}

func (m *ShopRepositoryMock) GetShopByID(ctx context.Context, id string) (*models.Shop, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Shop), args.Error(1)
}

func (m *ShopRepositoryMock) CreateShop(ctx context.Context, shop *models.Shop) error {
	args := m.Called(shop)
	return args.Error(0)
}

func (m *ShopRepositoryMock) UpdateShop(ctx context.Context, shop *models.Shop) error {
	args := m.Called(shop)
	return args.Error(0)
}

func (m *ShopRepositoryMock) DeleteShop(ctx context.Context, id string) error {
	args := m.Called(id)
	return args.Error(0)
}
