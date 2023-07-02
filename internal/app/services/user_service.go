package services

import (
	"fmt"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/repositories"
)

type UserService struct {
	UserRepo repositories.UserRepository
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.UserRepo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return fmt.Errorf("name and email are required")
	}

	return s.UserRepo.SaveUser(user)
}
