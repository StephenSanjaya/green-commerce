package cmd

import (
	"crypto/tls"
	"log"
	"ms-gateaway/controller"
	pb "ms-gateaway/pb/user"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func UserClientGRPC() controller.UserControllerI {
	port := os.Getenv("PORT_GRPC")
	host := os.Getenv("MS_USER_HOST")
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	userConn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		log.Fatal(err)
	}

	userClient := pb.NewUserServiceClient(userConn)
	// defer authConn.Close()

	userCtrler := controller.NewUserController(userClient)

	return userCtrler
}
