package resolver

import (
	"context"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/app/services"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models/mutations"
)

type UserResolver struct {
	UserService *services.UserService
}

func NewUserResolver(userService *services.UserService) *UserResolver {
	return &UserResolver{
		UserService: userService,
	}
}

func (r *UserResolver) CreateUserMutation(ctx context.Context, input mutations.CreateUserInput) (*models.User, error) {
	user := &models.User{
		Username: input.Username,
		Email:    input.Email,
	}

	err := r.UserService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Tel:      user.Tel,
		Role:     user.Role,
	}, nil
}

func (r *UserResolver) GetUserQuery(ctx context.Context, id string) (*models.User, error) {
	user, err := r.UserService.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
