package graph

import (
	"context"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/graph/generated"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
)

func (r *queryResolver) GetUser(ctx context.Context, id string) (*models.User, error) {
	return r.UserService.GetUser(id)
}

func (r *queryResolver) GetShop(ctx context.Context, id string) (*models.Shop, error) {
	return r.ShopService.GetShopByID(id)
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
