package controller

import "github.com/labstack/echo/v4"

type AuthControllerI interface {
	RegisterAuth(c echo.Context) error
	LoginAuth(c echo.Context) error
}
