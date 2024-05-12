package cmd

import (
	"fmt"
	"log"
	"ms-auth/controller"
	pb "ms-auth/pb/auth"
	"net"

	"google.golang.org/grpc"
)

func InitGrpc(paymentCtrler *controller.AuthControllerImpl) {

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, paymentCtrler)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println(err)
	}

	log.Printf("server listening at %v", listen.Addr())
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("INI GRPC")
}
