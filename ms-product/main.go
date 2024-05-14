package main

import (
	"ms-product/cmd"
	"ms-product/config"
	"ms-product/controller"
	"ms-product/repository"
)

func main() {
	db := config.GetConnection()

	productRepo := repository.NewProductRepositoryImpl(db)
	productCtrler := controller.NewProductControllerImpl(productRepo)

	cmd.InitGrpc(productCtrler.(*controller.ProductControllerImpl))
}
