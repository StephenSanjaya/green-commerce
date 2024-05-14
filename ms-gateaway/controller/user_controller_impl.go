package controller

import (
	"context"
	pb "ms-gateaway/pb/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	userGRPC pb.UserServiceClient
}

func NewUserController(userGRPC pb.UserServiceClient) *UserControllerImpl {
	return &UserControllerImpl{userGRPC: userGRPC}
}

func (uc *UserControllerImpl) AddProductToCart(c echo.Context) error {
	var cartRequest pb.CartRequest
	if err := c.Bind(&cartRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	_, err := uc.userGRPC.AddProductToCart(context.Background(), &cartRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add product to cart"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Product added to cart successfully"})
}
