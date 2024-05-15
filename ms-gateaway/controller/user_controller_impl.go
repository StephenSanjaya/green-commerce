package controller

import (
	pb "ms-gateaway/pb/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
)

type UserControllerImpl struct {
	userGRPC pb.UserServiceClient
}

func NewUserController(userGRPC pb.UserServiceClient) UserControllerI {
	return &UserControllerImpl{userGRPC: userGRPC}
}

func (uc *UserControllerImpl) GetCartItems(c echo.Context) error {
	user_id := int(c.Get("id").(float64))
	req := &pb.GetCartItemsRequest{
		UserId: int64(user_id),
	}

	res, err := uc.userGRPC.GetCartItems(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "Failed to get products from cart: "+err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success get products from cart",
		"cart":    res.Carts,
	})
}

func (uc *UserControllerImpl) AddProductToCart(c echo.Context) error {
	var cartRequest pb.AddProductToCartRequest
	if err := c.Bind(&cartRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request: "+err.Error())
	}

	_, err := uc.userGRPC.AddProductToCart(c.Request().Context(), &cartRequest)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "Failed to add product to cart: "+err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Product added to cart successfully",
	})
}

func (uc *UserControllerImpl) TopUp(c echo.Context) error {
	req := &pb.TopUpRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request: "+err.Error())
	}

	_, err := uc.userGRPC.TopUp(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "failed to top up: "+err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "top up successful",
	})
}
