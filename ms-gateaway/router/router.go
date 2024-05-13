package router

import (
	"ms-gateaway/controller"

	"github.com/labstack/echo/v4"
)

type ControllerStruct struct {
	AuthCtrler    controller.AuthControllerI
	ProductCtrler controller.ProductControllerI
}

func (cs ControllerStruct) SetupRouter(e *echo.Echo) {

	auth := e.Group("/api/v1")
	{
		auth.POST("/register", cs.AuthCtrler.RegisterAuth)
		auth.POST("/login", cs.AuthCtrler.LoginAuth)
	}

	//product
	product := e.Group("/api/v1/products")
	{
		product.POST("", cs.ProductCtrler.AddProduct)
		product.GET("", cs.ProductCtrler.GetAllProduct)
		product.GET("/:name", cs.ProductCtrler.GetProduct)
		product.PUT("/:id", cs.ProductCtrler.UpdateProduct)
		product.DELETE("/:id", cs.ProductCtrler.DeleteProduct)
	}

}
