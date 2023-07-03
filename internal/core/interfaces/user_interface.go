package interfaces

import "github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"

type UserInterface interface {
	GetAllUsers() ([]*models.User, error)
	GetUserByID(id string) (*models.User, error)
	SaveUser(user *models.User) error
}

type UserServiceInterface interface {
	GetAllUsers() ([]*models.User, error)
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
}
