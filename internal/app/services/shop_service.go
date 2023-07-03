package services

import (
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/interfaces"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
)

type ShopService struct {
	ShopRepo interfaces.ShopInterface
}

func NewShopService(shopRepo interfaces.ShopInterface) *ShopService {
	return &ShopService{
		ShopRepo: shopRepo,
	}
}

func (s *ShopService) CreateShop(shop *models.Shop) error {
	return s.ShopRepo.CreateShop(shop)
}

func (s *ShopService) GetShopByID(id string) (*models.Shop, error) {
	return s.ShopRepo.GetShopByID(id)
}
