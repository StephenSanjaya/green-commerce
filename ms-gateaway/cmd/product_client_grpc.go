package cmd

import (
	"log"
	"ms-gateaway/controller"
	pb "ms-gateaway/pb/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ProductClientGRPC() controller.ProductControllerI {
	productConn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	productClient := pb.NewProductServiceClient(productConn)
	// defer authConn.Close()

	productCtrler := controller.NewProductController(productClient)

	return productCtrler
}
