package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"ms-gateaway/controller"
	pb "ms-gateaway/pb/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockUserServiceClient struct {
	mock.Mock
}

func (m *MockUserServiceClient) AddProductToCart(ctx context.Context, in *pb.CartRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (m *MockUserServiceClient) TopUp(ctx context.Context, in *pb.TopUpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

// Define a struct without the internal mutex for JSON marshalling
type CartRequest struct {
	ProductId     int64   `json:"product_id"`
	Quantity      int64   `json:"quantity"`
	SubTotalPrice float64 `json:"sub_total_price"`
}

type TopUpRequest struct {
	Amount float64 `json:"amount"`
}

func TestAddProductToCart(t *testing.T) {
	mockUserService := new(MockUserServiceClient)
	controller := controller.NewUserController(mockUserService)
	e := echo.New()

	cartRequest := CartRequest{
		ProductId:     1,
		Quantity:      2,
		SubTotalPrice: 100.0,
	}

	mockUserService.On("AddProductToCart", mock.Anything, &pb.CartRequest{
		ProductId:     cartRequest.ProductId,
		Quantity:      cartRequest.Quantity,
		SubTotalPrice: cartRequest.SubTotalPrice,
	}).Return(&emptypb.Empty{}, nil)

	reqBody, _ := json.Marshal(cartRequest)
	req := httptest.NewRequest(http.MethodPost, "/cart", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controller.AddProductToCart(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "Product added to cart successfully", response["message"])
	}
}

func TestTopUp(t *testing.T) {
	mockUserService := new(MockUserServiceClient)
	controller := controller.NewUserController(mockUserService)
	e := echo.New()

	topUpRequest := TopUpRequest{
		Amount: 50.0,
	}

	mockUserService.On("TopUp", mock.Anything, &pb.TopUpRequest{
		Amount: topUpRequest.Amount,
	}).Return(&emptypb.Empty{}, nil)

	reqBody, _ := json.Marshal(topUpRequest)
	req := httptest.NewRequest(http.MethodPost, "/topup", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controller.TopUp(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "top up successful", response["message"])
	}
}
