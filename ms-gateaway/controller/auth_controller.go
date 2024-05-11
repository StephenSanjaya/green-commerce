package controller

import "github.com/labstack/echo/v4"

type AuthControllerI interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}
