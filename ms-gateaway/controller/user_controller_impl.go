package controller

import (
	"ms-gateaway/helper"
	pb "ms-gateaway/pb/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
)

// UserControllerImpl handles user-related operations
type UserControllerImpl struct {
	userGRPC pb.UserServiceClient
}

// NewUserController creates a new UserController
func NewUserController(userGRPC pb.UserServiceClient) UserControllerI {
	return &UserControllerImpl{userGRPC: userGRPC}
}

// GetCartItems godoc
// @Summary Get cart items
// @Description Get all items in the user's cart
// @Tags users
// @Produce json
// @Success 200 {object} map[string]interface{} "success get products from cart"
// @Failure 500 {object} helper.HTTPError "failed to get products from cart"
// @Router /cart [get]
func (uc *UserControllerImpl) GetCartItems(c echo.Context) error {
	user_id := int(c.Get("id").(float64))
	req := &pb.GetCartItemsRequest{
		UserId: int64(user_id),
	}

	res, err := uc.userGRPC.GetCartItems(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "Failed to get products from cart: "+err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get products from cart",
		"cart":    res.Carts,
	})
}

// AddProductToCart godoc
// @Summary Add product to cart
// @Description Add a product to the user's cart
// @Tags users
// @Accept json
// @Produce json
// @Param cartRequest body pb.AddProductToCartRequest true "Cart Request"
// @Success 201 {object} map[string]interface{} "Product added to cart successfully"
// @Failure 400 {object} helper.HTTPError "invalid body request"
// @Failure 500 {object} helper.HTTPError "failed to add product to cart"
// @Router /cart [post]
func (uc *UserControllerImpl) AddProductToCart(c echo.Context) error {
	user_id := int(c.Get("id").(float64))
	var cartRequest pb.AddProductToCartRequest
	if err := c.Bind(&cartRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request: "+err.Error())
	}
	cartRequest.UserId = int64(user_id)

	_, err := uc.userGRPC.AddProductToCart(c.Request().Context(), &cartRequest)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "Failed to add product to cart: "+err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Product added to cart successfully",
	})
}

// TopUp godoc
// @Summary Top up user balance
// @Description Top up the user's balance
// @Tags users
// @Accept json
// @Produce json
// @Param topUpRequest body pb.TopUpRequest true "Top Up Request"
// @Success 200 {object} map[string]interface{} "top up successful"
// @Failure 400 {object} helper.HTTPError "invalid body request"
// @Failure 500 {object} helper.HTTPError "failed to top up"
// @Router /top-up [post]
func (uc *UserControllerImpl) TopUp(c echo.Context) error {
	user_id := int(c.Get("id").(float64))
	email := c.Get("email").(string)
	req := &pb.TopUpRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request: "+err.Error())
	}
	req.UserId = int64(user_id)

	_, err := uc.userGRPC.TopUp(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "failed to top up: "+err.Error())
	}

	url, errInv := helper.CreateInvoiceTopUp(req.Amount)
	if errInv != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create invoice: "+errInv.Error())
	}

	helper.SendSuccessCheckout(email, url)

	return c.JSON(http.StatusOK, echo.Map{
		"message":     "top up successful",
		"invoice_url": url,
	})
}