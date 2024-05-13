package controller

import "github.com/labstack/echo/v4"

type ProductControllerI interface {
	GetAllProduct(c echo.Context) error
	GetProduct(c echo.Context) error
	AddProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
}
