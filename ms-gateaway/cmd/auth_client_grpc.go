package cmd

import (
	"crypto/tls"
	"log"
	"ms-gateaway/controller"
	pb "ms-gateaway/pb/auth"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func AuthClientGRPC() controller.AuthControllerI {
	port := os.Getenv("PORT_GRPC")
	host := os.Getenv("MS_AUTH_HOST")
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	authConn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		log.Fatal(err)
	}

	authClient := pb.NewAuthServiceClient(authConn)
	// defer authConn.Close()

	authCtrler := controller.NewAuthController(authClient)

	return authCtrler
}
