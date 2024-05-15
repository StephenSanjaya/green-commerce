package main

import (
	"ms-user/cmd"
	"ms-user/config"
	"ms-user/controller"
	"ms-user/repository"
)

func main() {
	db := config.GetConnection()

	userRepo := repository.NewUserRepositoryImpl(db)
	userCtrler := controller.NewUserControllerImpl(userRepo)

	cmd.InitGrpc(userCtrler.(*controller.UserControllerImpl))
}
