package main

import (
	"ms-auth/cmd"
	"ms-auth/config"
	"ms-auth/controller"
	"ms-auth/repository"
)

func main() {
	db := config.GetConnection()

	authRepo := repository.NewAuthRepositoryImpl(db)
	authCtrler := controller.NewAuthControllerImpl(authRepo)

	cmd.InitGrpc(authCtrler.(*controller.AuthControllerImpl))
}
