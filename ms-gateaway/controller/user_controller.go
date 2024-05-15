package controller

import "github.com/labstack/echo/v4"

type UserControllerI interface {
	AddProductToCart(c echo.Context) error
	TopUp(c echo.Context) error
	GetCartItems(c echo.Context) error
}
