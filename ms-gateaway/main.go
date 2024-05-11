package main

import (
	"ms-gateaway/cmd"
	"ms-gateaway/middleware"
	"ms-gateaway/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.MiddlewareLogging)
	e.HTTPErrorHandler = middleware.ErrorHandler

	authCtrler := cmd.AuthClientGRPC()

	ctrlers := &router.ControllerStruct{
		AuthCtrler: authCtrler,
	}

	ctrlers.SetupRouter(e)

	e.Logger.Fatal(e.Start(":8081"))
}
