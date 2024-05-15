package controller

import "github.com/labstack/echo/v4"

type OrderControllerI interface {
	CheckoutOrder(c echo.Context) error
	PayOrder(c echo.Context) error
}
