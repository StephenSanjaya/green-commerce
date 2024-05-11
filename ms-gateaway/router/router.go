package router

import (
	"ms-gateaway/controller"

	"github.com/labstack/echo/v4"
)

type ControllerStruct struct {
	AuthCtrler controller.AuthControllerI
}

func (cs ControllerStruct) SetupRouter(e *echo.Echo) {

	auth := e.Group("/api/v1")
	{
		auth.POST("/register", cs.AuthCtrler.Register)
		auth.POST("/login", cs.AuthCtrler.Login)
	}

}
