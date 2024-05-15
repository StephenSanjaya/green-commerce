package main

import (
	"ms-gateaway/cmd"
	"ms-gateaway/middleware"
	"ms-gateaway/router"
	"os"

	_ "ms-gateaway/docs"

	"github.com/labstack/echo/v4"
)

// @title Auth Service API
// @version 1.0
// @description This is an auth service.
// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()
	e.Use(middleware.MiddlewareLogging)
	e.HTTPErrorHandler = middleware.ErrorHandler

	authCtrler := cmd.AuthClientGRPC()
	productCtrler := cmd.ProductClientGRPC()
	userCtrler := cmd.UserClientGRPC()
	orderCtrler := cmd.OrderClientGRPC()

	ctrlers := &router.ControllerStruct{
		AuthCtrler:    authCtrler,
		ProductCtrler: productCtrler,
		UserCtrler:    userCtrler,
		OrderCtrler:   orderCtrler,
	}

	ctrlers.SetupRouter(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
