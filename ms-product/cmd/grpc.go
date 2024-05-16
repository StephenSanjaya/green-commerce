package cmd

import (
	"fmt"
	"log"
	"ms-product/controller"
	pb "ms-product/pb/product"
	"net"
	"os"

	"google.golang.org/grpc"
)

func InitGrpc(productCtrler *controller.ProductControllerImpl) {

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, productCtrler)

	port := os.Getenv("PORT")
	listen, err := net.Listen("tcp", ":"+port)
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
