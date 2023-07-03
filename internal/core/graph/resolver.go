package graph

import (
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/app/services"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/interfaces"
)

type Resolver struct {
	UserService *services.UserService
	ShopService *services.ShopService
}

func NewResolver(userRepo interfaces.UserInterface, shopRepo interfaces.ShopInterface) *Resolver {
	return &Resolver{
		UserService: services.NewUserService(userRepo),
		ShopService: services.NewShopService(shopRepo),
	}
}
