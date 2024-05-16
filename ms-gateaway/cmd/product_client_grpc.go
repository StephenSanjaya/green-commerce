package cmd

import (
	"crypto/tls"
	"log"
	"ms-gateaway/controller"
	pb "ms-gateaway/pb/product"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func ProductClientGRPC() controller.ProductControllerI {
	port := os.Getenv("PORT_GRPC")
	host := os.Getenv("MS_PRODUCT_HOST")
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	productConn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		log.Fatal(err)
	}

	productClient := pb.NewProductServiceClient(productConn)
	// defer authConn.Close()

	productCtrler := controller.NewProductController(productClient)

	return productCtrler
}
