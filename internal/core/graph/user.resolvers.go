package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/app/resolver"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/graph/generated"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/graph/models_gen"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
)

type UserResolver struct {
	*resolver.UserResolver
}

func NewUserResolver(userService *resolver.UserResolver) *UserResolver {
	return &UserResolver{
		UserResolver: userService,
	}
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models_gen.CreateUserInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*models.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// Role is the resolver for the role field.
func (r *userResolver) Role(ctx context.Context, obj *models.User) (models_gen.Role, error) {
	panic(fmt.Errorf("not implemented: Role - role"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
