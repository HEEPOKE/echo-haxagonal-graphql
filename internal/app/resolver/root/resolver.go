package root

import (
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/app/resolver"
)

type RootResolver struct {
	UserResolver *resolver.UserResolver
	ShopResolver *resolver.ShopResolver
}

func NewRootResolver(userResolver *resolver.UserResolver, shopResolver *resolver.ShopResolver) *RootResolver {
	return &RootResolver{
		UserResolver: userResolver,
		ShopResolver: shopResolver,
	}
}
