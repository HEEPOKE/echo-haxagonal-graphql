package resolver

import (
	"context"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/interfaces"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
)

type ShopResolver struct {
	ShopService interfaces.ShopServiceInterface
}

func NewShopResolver(shopService interfaces.ShopServiceInterface) *ShopResolver {
	return &ShopResolver{
		ShopService: shopService,
	}
}

// func (r *ShopResolver) CreateShop(ctx context.Context, input models_gen.CreateShopInput) (*models.Shop, error) {
// 	shop := &models.Shop{
// 		Name:    input.Name,
// 		Address: input.Address,
// 	}

// 	err := r.ShopService.CreateShop(shop)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return shop, nil
// }

func (r *ShopResolver) GetShop(ctx context.Context, id string) (*models.Shop, error) {
	shop, err := r.ShopService.GetShopByID(id)
	if err != nil {
		return nil, err
	}

	return shop, nil
}
