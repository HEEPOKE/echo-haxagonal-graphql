package interfaces

import "github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"

type ShopInterface interface {
	GetAllShops() ([]*models.Shop, error)
	GetShopByID(id string) (*models.Shop, error)
	CreateShop(shop *models.Shop) error
}

type ShopServiceInterface interface {
	GetAllShops() ([]*models.Shop, error)
	GetShopByID(id string) (*models.Shop, error)
	CreateShop(shop *models.Shop) error
}
