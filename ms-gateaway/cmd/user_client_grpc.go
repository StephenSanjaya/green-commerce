package cmd

import (
	"log"
	"ms-gateaway/controller"
	pb "ms-gateaway/pb/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClientGRPC() controller.UserControllerI {
	userConn, err := grpc.Dial(":50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	userClient := pb.NewUserServiceClient(userConn)
	// defer authConn.Close()

	userCtrler := controller.NewUserController(userClient)

	return userCtrler
}
