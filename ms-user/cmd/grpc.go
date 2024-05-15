package cmd

import (
	"fmt"
	"log"
	"ms-user/controller"
	pb "ms-user/pb/user"
	"net"

	"google.golang.org/grpc"
)

func InitGrpc(userCtrler *controller.UserControllerImpl) {

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userCtrler)

	listen, err := net.Listen("tcp", ":50053")
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
