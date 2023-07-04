package mocks

import (
	"context"

	"github.com/HEEPOKE/echo-haxagonal-graphql/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	args := m.Called()
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserRepositoryMock) SaveUser(ctx context.Context, user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) UpdateUser(ctx context.Context, user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) DeleteUser(ctx context.Context, id string) error {
	args := m.Called(id)
	return args.Error(0)
}
