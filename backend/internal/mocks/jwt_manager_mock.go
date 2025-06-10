package mocks

import (
	"go-task-manager-backend/internal/auth"

	"github.com/stretchr/testify/mock"
)

type MockJWTManager struct {
	mock.Mock
}

func (m *MockJWTManager) GenerateToken(userID int) (string, error) {
	args := m.Called(userID)
	return args.String(0), args.Error(1)
}

func (m *MockJWTManager) ParseToken(tokenStr string) (*auth.Claims, error) {
	args := m.Called(tokenStr)
	return args.Get(0).(*auth.Claims), args.Error(1)
}
