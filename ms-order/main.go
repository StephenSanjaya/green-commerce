package main

import (
	"ms-order/cmd"
	"ms-order/config"
	"ms-order/controller"
	"ms-order/repository"
)

func main() {
	mongo := config.GetMongoConnection().Orders
	psql := config.GetPsqlConnection()

	orderRepo := repository.NewOrderRepositoryImpl(mongo, psql)
	orderCtrler := controller.NewOrderControllerImpl(orderRepo)

	cmd.InitGrpc(orderCtrler.(*controller.OrderControllerImpl))
}
