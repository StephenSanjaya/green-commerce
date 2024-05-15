package controller

import (
	"context"
	pb "ms-gateaway/pb/product"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductControllerImpl struct {
	productGRPC pb.ProductServiceClient
}

func NewProductController(productGRPC pb.ProductServiceClient) ProductControllerI {
	return &ProductControllerImpl{productGRPC: productGRPC}
}

// GetAllProduct godoc
// @Summary Get all products
// @Description Get all products from the database
// @Tags products
// @Produce json
// @Success 200 {object} map[string]interface{} "success get all products"
// @Failure 500 {object} helper.HTTPError
// @Router /products [get]
func (pc *ProductControllerImpl) GetAllProduct(c echo.Context) error {
	res, err := pc.productGRPC.GetAllProduct(c.Request().Context(), &emptypb.Empty{})
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "failed to get all products: "+err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  "success get all products",
		"products": res.Products,
	})
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get a product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{} "success get product"
// @Failure 400 {object} helper.HTTPError
// @Failure 404 {object} helper.HTTPError
// @Router /products/{id} [get]
func (pc *ProductControllerImpl) GetProduct(c echo.Context) error {
	id := c.Param("id")
	product_id, _ := strconv.Atoi(id)
	req := &pb.ProductId{
		ProductId: int64(product_id),
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request: "+err.Error())
	}

	res, err := pc.productGRPC.GetProduct(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "failed to get product: "+err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get product",
		"product": res,
	})
}

// AddProduct godoc
// @Summary Add a new product
// @Description Add a new product to the database
// @Tags products
// @Accept json
// @Produce json
// @Param product body pb.ProductRequest true "Product Request"
// @Success 201 {object} map[string]interface{} "success create product"
// @Failure 400 {object} helper.HTTPError
// @Failure 500 {object} helper.HTTPError
// @Router /products [post]
func (pc *ProductControllerImpl) AddProduct(c echo.Context) error {
	req := &pb.ProductRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request: "+err.Error())
	}

	res, err := pc.productGRPC.AddProduct(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "failed to add product: "+err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success create product",
		"product": res,
	})
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Delete a product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{} "success delete product"
// @Failure 400 {object} helper.HTTPError
// @Failure 404 {object} helper.HTTPError
// @Router /products/{id} [delete]
func (pc *ProductControllerImpl) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	product_id, _ := strconv.Atoi(id)
	req := &pb.ProductId{
		ProductId: int64(product_id),
	}

	_, err := pc.productGRPC.DeleteProduct(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "failed to delete product: "+err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete product",
	})
}

// UpdateProduct godoc
// @Summary Update a product by ID
// @Description Update a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body pb.ProductRequest true "Product Request"
// @Success 200 {object} map[string]interface{} "success update product"
// @Failure 400 {object} helper.HTTPError
// @Failure 404 {object} helper.HTTPError
// @Router /products/{id} [put]
func (pc *ProductControllerImpl) UpdateProduct(c echo.Context) error {
	id := c.Param("id")

	req := &pb.ProductRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request: "+err.Error())
	}

	md := metadata.New(map[string]string{
		"product_id": id,
	})
	ctx := context.Background()
	ctxWithMetadata := metadata.NewOutgoingContext(ctx, md)
	res, err := pc.productGRPC.UpdateProduct(ctxWithMetadata, req)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "failed to update product: "+err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update product",
		"product": res,
	})
}
