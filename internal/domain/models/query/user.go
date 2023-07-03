package query

import (
	"time"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
)

type UserQuery struct {
	ID        string
	Username  string
	Email     string
	Password  string
	Tel       string
	Role      models.Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
