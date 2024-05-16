package cmd

import (
	"crypto/tls"
	"log"
	"ms-gateaway/controller"
	pb "ms-gateaway/pb/order"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func OrderClientGRPC() controller.OrderControllerI {
	port := os.Getenv("PORT_GRPC")
	host := os.Getenv("MS_ORDER_HOST")
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	orderConn, err := grpc.Dial(host+":"+port, opts...)
	if err != nil {
		log.Fatal(err)
	}
	orderClient := pb.NewOrderServiceClient(orderConn)
	// defer authConn.Close()

	orderCtrler := controller.NewOrderControllerImpl(orderClient)

	return orderCtrler
}
