package main

import (
	"ms-gateaway/cmd"
	"ms-gateaway/middleware"
	"ms-gateaway/router"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.MiddlewareLogging)
	e.HTTPErrorHandler = middleware.ErrorHandler

	authCtrler := cmd.AuthClientGRPC()
	productCtrler := cmd.ProductClientGRPC()
	userCtrler := cmd.UserClientGRPC()

	ctrlers := &router.ControllerStruct{
		AuthCtrler:    authCtrler,
		ProductCtrler: productCtrler,
		UserCtrler:    userCtrler,
	}

	ctrlers.SetupRouter(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
