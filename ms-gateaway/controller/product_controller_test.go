package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"ms-gateaway/controller"
	pb "ms-gateaway/pb/product"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockProductServiceClient struct {
	mock.Mock
}

func (m *MockProductServiceClient) GetAllProduct(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pb.ProductResponses, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.ProductResponses), args.Error(1)
}

func (m *MockProductServiceClient) GetProduct(ctx context.Context, in *pb.ProductId, opts ...grpc.CallOption) (*pb.ProductResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.ProductResponse), args.Error(1)
}

func (m *MockProductServiceClient) AddProduct(ctx context.Context, in *pb.ProductRequest, opts ...grpc.CallOption) (*pb.ProductResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.ProductResponse), args.Error(1)
}

func (m *MockProductServiceClient) DeleteProduct(ctx context.Context, in *pb.ProductId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (m *MockProductServiceClient) UpdateProduct(ctx context.Context, in *pb.ProductRequest, opts ...grpc.CallOption) (*pb.ProductResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.ProductResponse), args.Error(1)
}

func TestGetAllProduct(t *testing.T) {
	mockProductService := new(MockProductServiceClient)
	controller := controller.NewProductController(mockProductService)
	e := echo.New()

	productResponses := &pb.ProductResponses{
		Products: []*pb.ProductResponse{
			{ProductId: 1, CategoryId: 1, Name: "Product 1", Description: "Description 1", Stock: 10, Price: 100.0},
			{ProductId: 2, CategoryId: 2, Name: "Product 2", Description: "Description 2", Stock: 20, Price: 200.0},
		},
	}

	mockProductService.On("GetAllProduct", mock.Anything, &emptypb.Empty{}).Return(productResponses, nil)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controller.GetAllProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success get all products", response["message"])
		assert.NotNil(t, response["products"])
	}
}

func TestGetProduct(t *testing.T) {
	mockProductService := new(MockProductServiceClient)
	controller := controller.NewProductController(mockProductService)
	e := echo.New()

	productResponse := &pb.ProductResponse{
		ProductId:   1,
		CategoryId:  1,
		Name:        "Product 1",
		Description: "Description 1",
		Stock:       10,
		Price:       100.0,
	}

	mockProductService.On("GetProduct", mock.Anything, &pb.ProductId{ProductId: 1}).Return(productResponse, nil)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, controller.GetProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "sucess get product", response["message"])
		assert.NotNil(t, response["product"])
	}
}

func TestAddProduct(t *testing.T) {
	mockProductService := new(MockProductServiceClient)
	controller := controller.NewProductController(mockProductService)
	e := echo.New()

	productRequest := &pb.ProductRequest{
		CategoryId:  1,
		Name:        "Product 1",
		Description: "Description 1",
		Stock:       10,
		Price:       100.0,
	}

	productResponse := &pb.ProductResponse{
		ProductId:   1,
		CategoryId:  1,
		Name:        "Product 1",
		Description: "Description 1",
		Stock:       10,
		Price:       100.0,
	}

	mockProductService.On("AddProduct", mock.Anything, productRequest).Return(productResponse, nil)

	reqBody, _ := json.Marshal(productRequest)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controller.AddProduct(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success create product", response["message"])
		assert.NotNil(t, response["product"])
	}
}

func TestDeleteProduct(t *testing.T) {
	mockProductService := new(MockProductServiceClient)
	controller := controller.NewProductController(mockProductService)
	e := echo.New()

	mockProductService.On("DeleteProduct", mock.Anything, &pb.ProductId{ProductId: 1}).Return(&emptypb.Empty{}, nil)

	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, controller.DeleteProduct(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success delete product", response["message"])
	}
}

func TestUpdateProduct(t *testing.T) {
	mockProductService := new(MockProductServiceClient)
	controller := controller.NewProductController(mockProductService)
	e := echo.New()

	productRequest := &pb.ProductRequest{
		CategoryId:  1,
		Name:        "Updated Product",
		Description: "Updated Description",
		Stock:       15,
		Price:       150.0,
	}

	productResponse := &pb.ProductResponse{
		ProductId:   1,
		CategoryId:  1,
		Name:        "Updated Product",
		Description: "Updated Description",
		Stock:       15,
		Price:       150.0,
	}

	mockProductService.On("UpdateProduct", mock.Anything, productRequest).Return(productResponse, nil)

	reqBody, _ := json.Marshal(productRequest)
	req := httptest.NewRequest(http.MethodPut, "/products/1", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, controller.UpdateProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "success update product", response["message"])
		assert.NotNil(t, response["product"])
	}
}
