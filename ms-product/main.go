package main

import (
	"ms-product/cmd"
	"ms-product/config"
	"ms-product/controller"
	"ms-product/repository"
)

func main() {
	db := config.GetConnection()
	redis := config.SetupRedis()

	productRepo := repository.NewProductRepositoryImpl(db, redis)
	productCtrler := controller.NewProductControllerImpl(productRepo)

	cmd.InitGrpc(productCtrler.(*controller.ProductControllerImpl))
}
