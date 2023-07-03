package root

import (
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/app/resolver"
)

type RootResolver struct {
	*resolver.UserResolver
	*resolver.ShopResolver
}

func NewRootResolver(userResolver *resolver.UserResolver, shopResolver *resolver.ShopResolver) *RootResolver {
	return &RootResolver{
		UserResolver: userResolver,
		ShopResolver: shopResolver,
	}
}
