package cmd

import (
	"log"
	"ms-gateaway/controller"
	pb "ms-gateaway/pb/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AuthClientGRPC() controller.AuthControllerI {
	authConn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	authClient := pb.NewAuthServiceClient(authConn)
	// defer authConn.Close()

	authCtrler := controller.NewAuthController(authClient)

	return authCtrler
}
