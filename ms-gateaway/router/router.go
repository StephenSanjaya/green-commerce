package router

import (
	"ms-gateaway/controller"
	"ms-gateaway/middleware"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type ControllerStruct struct {
	AuthCtrler    controller.AuthControllerI
	ProductCtrler controller.ProductControllerI
	UserCtrler    controller.UserControllerI
}

func (cs ControllerStruct) SetupRouter(e *echo.Echo) {

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api := e.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", cs.AuthCtrler.RegisterAuth)
			auth.POST("/login", cs.AuthCtrler.LoginAuth)
		}

		product := api.Group("/products")
		{
			product.Use(middleware.AuthMiddleware("admin"))
			{
				product.POST("", cs.ProductCtrler.AddProduct)
				product.PUT("/:id", cs.ProductCtrler.UpdateProduct)
				product.DELETE("/:id", cs.ProductCtrler.DeleteProduct)
			}
			product.GET("", cs.ProductCtrler.GetAllProduct)
			product.GET("/:id", cs.ProductCtrler.GetProduct)
		}

		user := api.Group("/user")
		user.Use(middleware.AuthMiddleware("user"))
		{
			user.POST("/topup", cs.UserCtrler.TopUp)
			user.POST("/add-item", cs.UserCtrler.AddProductToCart)
		}
	}
}
