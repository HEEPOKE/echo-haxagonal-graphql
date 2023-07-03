package mutations

import "github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"

type CreateUserInput struct {
	UserName string
	Email    string
	Password string
	Tel      string
	Role     models.Role
}
