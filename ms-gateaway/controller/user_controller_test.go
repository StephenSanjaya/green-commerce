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

func (m *MockUserServiceClient) GetCartItems(ctx context.Context, in *pb.GetCartItemsRequest, opts ...grpc.CallOption) (*pb.GetCartItemsResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.GetCartItemsResponse), args.Error(1)
}

func (m *MockUserServiceClient) AddProductToCart(ctx context.Context, in *pb.AddProductToCartRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	return &emptypb.Empty{}, args.Error(1)
}

func (m *MockUserServiceClient) TopUp(ctx context.Context, in *pb.TopUpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	return &emptypb.Empty{}, args.Error(1)
}

type CartRequest struct {
	UserId        int64   `json:"user_id"`
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
		UserId:        1,
		ProductId:     1,
		Quantity:      2,
		SubTotalPrice: 100.0,
	}

	mockUserService.On("AddProductToCart", mock.Anything, &pb.AddProductToCartRequest{
		UserId:        cartRequest.UserId,
		ProductId:     cartRequest.ProductId,
		Quantity:      cartRequest.Quantity,
		SubTotalPrice: cartRequest.SubTotalPrice,
	}).Return(&emptypb.Empty{}, nil)

	reqBody, _ := json.Marshal(cartRequest)
	req := httptest.NewRequest(http.MethodPost, "/cart", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Setting user_id in context
	c.Set("id", float64(1))

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
		UserId: 1,
		Amount: topUpRequest.Amount,
	}).Return(&emptypb.Empty{}, nil)

	reqBody, _ := json.Marshal(topUpRequest)
	req := httptest.NewRequest(http.MethodPost, "/topup", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Setting user_id in context
	c.Set("id", float64(1))

	if assert.NoError(t, controller.TopUp(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "top up successful", response["message"])
	}
}

func TestGetCartItems(t *testing.T) {
	mockUserService := new(MockUserServiceClient)
	controller := controller.NewUserController(mockUserService)
	e := echo.New()

	mockResponse := &pb.GetCartItemsResponse{
		Carts: []*pb.Cart{
			{
				ProductId:     1,
				Quantity:      2,
				SubTotalPrice: 100.0,
			},
		},
	}

	mockUserService.On("GetCartItems", mock.Anything, &pb.GetCartItemsRequest{
		UserId: 1,
	}).Return(mockResponse, nil)

	req := httptest.NewRequest(http.MethodGet, "/cart", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Setting user_id in context
	c.Set("id", float64(1))

	if assert.NoError(t, controller.GetCartItems(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success get products from cart", response["message"])
		assert.NotNil(t, response["cart"])
	}
}
