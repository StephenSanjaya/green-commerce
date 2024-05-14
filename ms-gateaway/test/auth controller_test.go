package test

import (
	"bytes"
	"context"
	"encoding/json"
	"ms-gateaway/controller"
	"ms-gateaway/helper"
	pb "ms-gateaway/pb/auth"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type MockAuthServiceClient struct {
	mock.Mock
}

func (m *MockAuthServiceClient) RegisterAuth(ctx context.Context, req *pb.RegisterRequest, opts ...grpc.CallOption) (*pb.RegisterResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.RegisterResponse), args.Error(1)
}

func (m *MockAuthServiceClient) LoginAuth(ctx context.Context, req *pb.LoginRequest, opts ...grpc.CallOption) (*pb.LoginResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.LoginResponse), args.Error(1)
}

type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) CreateJWT(user *pb.LoginResponse) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func TestRegisterAuth(t *testing.T) {
	mockAuthService := new(MockAuthServiceClient)
	controller := controller.NewAuthController(mockAuthService)
	e := echo.New()
	registerRequest := &pb.RegisterRequest{
		FullName: "Test User",
		Email:    "testuser@example.com",
		Password: "testpass",
		Address:  "123 Test Street",
	}
	reqBody, _ := json.Marshal(registerRequest)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Update the expected request to match any hashed password
	expectedRequest := mock.MatchedBy(func(req *pb.RegisterRequest) bool {
		return req.FullName == "Test User" &&
			req.Email == "testuser@example.com" &&
			req.Address == "123 Test Street" &&
			len(req.Password) > 0 // Any non-empty hashed password
	})

	registerResponse := &pb.RegisterResponse{
		UserId:   123,
		FullName: "Test User",
		Email:    "testuser@example.com",
		Balance:  0.0,
		Address:  "123 Test Street",
		Role:     "user",
	}
	mockAuthService.On("RegisterAuth", mock.Anything, expectedRequest).Return(registerResponse, nil)

	if assert.NoError(t, controller.RegisterAuth(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success register", response["message"])
		assert.NotNil(t, response["user"])
	}
}

func TestLoginAuth(t *testing.T) {
	mockAuthService := new(MockAuthServiceClient)
	controller := controller.NewAuthController(mockAuthService)
	e := echo.New()
	loginRequest := &pb.LoginRequest{
		Email:    "testuser@example.com",
		Password: "testpass",
	}
	reqBody, _ := json.Marshal(loginRequest)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	hashedPassword := helper.HashedPassword("testpass")
	loginResponse := &pb.LoginResponse{
		UserId:   123,
		Email:    "testuser@example.com",
		Password: hashedPassword,
	}
	mockAuthService.On("LoginAuth", mock.Anything, loginRequest).Return(loginResponse, nil)

	os.Setenv("SECRET_KEY", "mysecretkey")
	token, _ := helper.CreateJWT(loginResponse)

	if assert.NoError(t, controller.LoginAuth(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success login", response["message"])
		assert.Equal(t, token, response["jwt"])
	}
}
