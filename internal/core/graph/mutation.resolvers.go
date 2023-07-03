package graph

import (
	"context"
	"time"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/graph/generated"
	model "github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/graph/models_gen"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	user := &models.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
		Tel:       input.Tel,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := r.UserService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) CreateShop(ctx context.Context, input model.CreateShopInput) (*models.Shop, error) {
	shop := &models.Shop{
		Name:      input.Name,
		Address:   input.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := r.ShopService.CreateShop(shop)
	if err != nil {
		return nil, err
	}

	return shop, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
