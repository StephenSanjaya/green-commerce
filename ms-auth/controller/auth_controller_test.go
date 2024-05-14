package controller

import (
	"context"
	"ms-auth/pb/auth"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) Insert(req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	args := m.Called(req)
	if res, ok := args.Get(0).(*auth.RegisterResponse); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthRepository) FindUser(req *auth.LoginRequest) (*auth.LoginResponse, error) {
	args := m.Called(req)
	if res, ok := args.Get(0).(*auth.LoginResponse); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestRegisterAuth(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	controller := NewAuthControllerImpl(mockRepo)

	request := &auth.RegisterRequest{
		FullName: "joni",
		Email:    "joni@gmail.com",
		Password: "password123",
		Address:  "garuda",
	}

	expectedResponse := &auth.RegisterResponse{
		UserId:   1,
		FullName: "joni",
		Email:    "joni@gmail.com",
		Balance:  100.0,
		Address:  "garuda",
		Role:     "user",
	}

	mockRepo.On("Insert", request).Return(expectedResponse, nil)

	result, err := controller.RegisterAuth(context.Background(), request)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)

	mockRepo.AssertExpectations(t)
}

func TestLoginAuth(t *testing.T) {
	mockRepo := new(MockAuthRepository)
	controller := NewAuthControllerImpl(mockRepo)

	request := &auth.LoginRequest{
		Email:    "joni@gmail.com",
		Password: "password123",
	}

	expectedResponse := &auth.LoginResponse{
		UserId: 1,
		Email:  "joni@gmail.com",
	}

	mockRepo.On("FindUser", request).Return(expectedResponse, nil)

	result, err := controller.LoginAuth(context.Background(), request)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)

	mockRepo.AssertExpectations(t)
}
