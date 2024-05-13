package controller

import (
	pb "ms-gateaway/pb/product"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductControllerImpl struct {
	productGRPC pb.ProductServiceClient
}

func NewProductController(productGRPC pb.ProductServiceClient) ProductControllerI {
	return &ProductControllerImpl{productGRPC: productGRPC}
}

func (pc *ProductControllerImpl) GetAllProduct(c echo.Context) error {
	res, err := pc.productGRPC.GetAllProduct(c.Request().Context(), &emptypb.Empty{})
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "get all product failed: "+err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":  "success get all products",
		"products": res,
	})
}

// GetProduct implements ProductControllerI.
func (pc *ProductControllerImpl) GetProduct(c echo.Context) error {
	req := &pb.ProductName{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request: "+err.Error())
	}
	res, err := pc.productGRPC.GetProduct(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "failed to get product: : "+err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "sucess get product",
		"product": res,
	})
}

// AddProduct implements ProductControllerI.
func (pc *ProductControllerImpl) AddProduct(c echo.Context) error {
	req := &pb.ProductRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request"+err.Error())
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

// DeleteProduct implements ProductControllerI.
func (pc *ProductControllerImpl) DeleteProduct(c echo.Context) error {
	productId := c.Param("id")
	req := &pb.ProductId{
		ProductId: productId,
	}
	res, err := pc.productGRPC.DeleteProduct(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "failed delete product: "+err.Error())
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success delete product",
		"product": res,
	})
}


// UpdateProduct implements ProductControllerI.
func (pc *ProductControllerImpl) UpdateProduct(c echo.Context) error {
	panic("unimplemented")
}
