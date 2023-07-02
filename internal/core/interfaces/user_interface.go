package interfaces

import "github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"

type UserRepository interface {
	GetUserByID(id string) (*models.User, error)
	SaveUser(user *models.User) error
}
