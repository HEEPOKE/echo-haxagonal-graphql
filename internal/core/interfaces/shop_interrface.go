package interfaces

import "github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"

type ShopInterface interface {
	GetShopByID(id string) (*models.Shop, error)
	CreateShop(shop *models.Shop) error
}

type ShopServiceInterface interface {
	GetShopByID(id string) (*models.Shop, error)
	CreateShop(shop *models.Shop) error
}
