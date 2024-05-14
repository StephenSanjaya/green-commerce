package router

import (
	"ms-gateaway/controller"

	"github.com/labstack/echo/v4"
)

type ControllerStruct struct {
	AuthCtrler    controller.AuthControllerI
	ProductCtrler controller.ProductControllerI
	UserCtrler    controller.UserControllerI
}

func (cs ControllerStruct) SetupRouter(e *echo.Echo) {

	api := e.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", cs.AuthCtrler.RegisterAuth)
			auth.POST("/login", cs.AuthCtrler.LoginAuth)
		}

		product := api.Group("/products")
		{
			product.POST("", cs.ProductCtrler.AddProduct)
			product.GET("", cs.ProductCtrler.GetAllProduct)
			product.GET("/:id", cs.ProductCtrler.GetProduct)
			product.PUT("/:id", cs.ProductCtrler.UpdateProduct)
			product.DELETE("/:id", cs.ProductCtrler.DeleteProduct)
		}

		user := api.Group("/user")
		{
			user.POST("/topup", cs.UserCtrler.TopUp)
			user.POST("/add-item", cs.UserCtrler.AddProductToCart)
		}
	}
}
