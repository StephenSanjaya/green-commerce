package cmd

import (
	"fmt"
	"log"
	"ms-product/controller"
	pb "ms-product/pb/product"
	"net"

	"google.golang.org/grpc"
)

func InitGrpc(productCtrler *controller.ProductControllerImpl) {

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, productCtrler)

	listen, err := net.Listen("tcp", ":50052")
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
