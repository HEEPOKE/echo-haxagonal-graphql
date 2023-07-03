package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/graph/generated"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
)

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*models.User, error) {
	return r.UserService.GetUser(id)
}

// GetShop is the resolver for the getShop field.
func (r *queryResolver) GetShop(ctx context.Context, id string) (*models.Shop, error) {
	return r.ShopService.GetShopByID(id)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
