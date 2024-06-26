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
	OrderCtrler   controller.OrderControllerI
}

func (cs ControllerStruct) SetupRouter(e *echo.Echo) {

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
			productAdmin := product.Group("")
			productAdmin.Use(middleware.AuthMiddleware("admin"))
			{
				productAdmin.POST("", cs.ProductCtrler.AddProduct)
				productAdmin.PUT("/:id", cs.ProductCtrler.UpdateProduct)
				productAdmin.DELETE("/:id", cs.ProductCtrler.DeleteProduct)
			}
			product.GET("", cs.ProductCtrler.GetAllProduct)
			product.GET("/:id", cs.ProductCtrler.GetProduct)
		}

		user := api.Group("/user")
		user.Use(middleware.AuthMiddleware("user"))
		{
			user.GET("/cart", cs.UserCtrler.GetCartItems)
			user.POST("/topup", cs.UserCtrler.TopUp)
			user.POST("/add-item", cs.UserCtrler.AddProductToCart)
		}

		order := api.Group("/orders")
		order.Use(middleware.AuthMiddleware("user"))
		{
			order.POST("/checkout", cs.OrderCtrler.CheckoutOrder)
			order.POST("/pay/:order_id", cs.OrderCtrler.PayOrder)
		}
	}
}
