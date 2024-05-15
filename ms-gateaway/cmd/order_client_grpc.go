package cmd

import (
	"log"
	"ms-gateaway/controller"
	pb "ms-gateaway/pb/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func OrderClientGRPC() controller.OrderControllerI {
	orderConn, err := grpc.Dial(":50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	orderClient := pb.NewOrderServiceClient(orderConn)
	// defer authConn.Close()

	orderCtrler := controller.NewOrderControllerImpl(orderClient)

	return orderCtrler
}
