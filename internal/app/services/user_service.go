package services

import (
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/core/interfaces"
	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
)

type UserService struct {
	UserRepo interfaces.UserInterface
}

func NewUserService(userRepo interfaces.UserInterface) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	return s.UserRepo.GetAllUsers()
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.UserRepo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepo.SaveUser(user)
}
