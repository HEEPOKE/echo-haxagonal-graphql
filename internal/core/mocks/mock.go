package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type DatabaseInterface interface {
	Collection(name string) CollectionInterface
}

type CollectionInterface interface {
	Find(ctx context.Context, filter interface{}, opts ...interface{}) (CursorInterface, error)
	InsertOne(ctx context.Context, document interface{}, opts ...interface{}) (interface{}, error)
}

type CursorInterface interface {
	Close(ctx context.Context) error
	Next(ctx context.Context) bool
	Decode(v interface{}) error
}

type MockDatabase struct {
	mock.Mock
}

type MockCollection struct {
	mock.Mock
}

type MockCursor struct {
	mock.Mock
}

func (m *MockDatabase) Collection(name string) CollectionInterface {
	args := m.Called(name)
	return args.Get(0).(CollectionInterface)
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...interface{}) (CursorInterface, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(CursorInterface), args.Error(1)
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...interface{}) (interface{}, error) {
	args := m.Called(ctx, document, opts)
	return args.Get(0), args.Error(1)
}

func (m *MockCursor) Close(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockCursor) Next(ctx context.Context) bool {
	args := m.Called(ctx)
	return args.Bool(0)
}

func (m *MockCursor) Decode(v interface{}) error {
	args := m.Called(v)
	return args.Error(0)
}
