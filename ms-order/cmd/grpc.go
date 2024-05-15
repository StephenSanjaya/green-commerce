package cmd

import (
	"fmt"
	"log"
	"ms-order/controller"
	pb "ms-order/pb/order"
	"net"

	"google.golang.org/grpc"
)

func InitGrpc(orderCtrler *controller.OrderControllerImpl) {

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderCtrler)

	listen, err := net.Listen("tcp", ":50054")
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
