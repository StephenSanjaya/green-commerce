package controller_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ms-gateaway/controller"
	pb "ms-gateaway/pb/order"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// MockOrderServiceClient is a mock implementation of the OrderServiceClient interface
type MockOrderServiceClient struct {
	mock.Mock
}

func (m *MockOrderServiceClient) PayOrder(ctx context.Context, in *pb.PayOrderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	if resp := args.Get(0); resp != nil {
		return resp.(*emptypb.Empty), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockOrderServiceClient) CheckoutOrder(ctx context.Context, in *pb.CheckoutOrderRequest, opts ...grpc.CallOption) (*pb.CheckoutOrderResponse, error) {
	args := m.Called(ctx, in)
	if resp := args.Get(0); resp != nil {
		return resp.(*pb.CheckoutOrderResponse), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestOrderController_CheckoutOrder(t *testing.T) {
	e := echo.New()
	mockClient := new(MockOrderServiceClient)
	orderController := controller.NewOrderControllerImpl(mockClient)

	t.Run("success", func(t *testing.T) {
		mockResponse := &pb.CheckoutOrderResponse{
			OrderId: "123",
			Products: []*pb.Product{
				{
					ProductId:     1,
					ProductName:   "Test Product",
					Quantity:      1,
					Price:         100.0,
					SubTotalPrice: 100.0,
				},
			},
			PaymentId:   1,
			Payment:     &pb.Payment{PaymentName: "Credit Card"},
			VoucherId:   1,
			Voucher:     &pb.Voucher{VoucherName: "Test Voucher"},
			TotalPrice:  100.0,
			OrderStatus: "Completed",
			OrderDate:   "2024-05-15",
		}
		// Mock the gRPC response for success case
		mockClient.On("CheckoutOrder", mock.Anything, mock.Anything).
			Return(mockResponse, nil)

		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/checkout", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("id", float64(1)) // Mock user ID
		c.Set("email", "test@email.com")

		// Call the controller method
		err := orderController.CheckoutOrder(c)
		require.NoError(t, err)

		// Assert the response
		assert.Equal(t, http.StatusCreated, rec.Code)
		var response map[string]interface{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "success checkout order", response["message"])
	})

	t.Run("failure", func(t *testing.T) {
		// Mock the gRPC error response for failure case
		mockClient.On("CheckoutOrder", mock.Anything, mock.Anything).Return(nil, status.Errorf(codes.Internal, "internal error"))
		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/checkout", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("id", float64(1)) // Mock user ID
		c.Set("email", "test@email.com")

		// Call the controller method
		// require.Error(t, err)

		// Assert the response
	})

}

func TestOrderController_PayOrder(t *testing.T) {
	e := echo.New()
	mockClient := new(MockOrderServiceClient)
	orderController := controller.NewOrderControllerImpl(mockClient)

	t.Run("success", func(t *testing.T) {
		// Mock the gRPC response for success case
		mockClient.On("PayOrder", mock.Anything, mock.Anything).
			Return(&emptypb.Empty{}, nil)

		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/pay/123", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("id", float64(1)) // Mock user ID
		c.SetParamNames("order_id")
		c.SetParamValues("123")

		// Call the controller method
		err := orderController.PayOrder(c)
		require.NoError(t, err)

		// Assert the response
		assert.Equal(t, http.StatusOK, rec.Code)
		var response map[string]interface{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "success pay order", response["message"])
	})

	t.Run("failure", func(t *testing.T) {
		// Mock the gRPC error response for failure case
		mockClient.On("PayOrder", mock.Anything, mock.Anything).
			Return(nil, status.Errorf(codes.Internal, "internal error"))

		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/pay/123", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("id", float64(1)) // Mock user ID
		c.SetParamNames("order_id")
		c.SetParamValues("123")

	})
}
