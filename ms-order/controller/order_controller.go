package controller

import (
	"context"
	pb "ms-order/pb/order"

	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderControllerI interface {
	CheckoutOrder(context.Context, *pb.CheckoutOrderRequest) (*pb.CheckoutOrderResponse, error)
	PayOrder(context.Context, *pb.PayOrderRequest) (*emptypb.Empty, error)
}
