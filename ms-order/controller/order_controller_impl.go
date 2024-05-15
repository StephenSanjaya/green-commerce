package controller

import (
	"context"
	"ms-order/pb/order"
	pb "ms-order/pb/order"
	"ms-order/repository"

	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderControllerImpl struct {
	pb.UnimplementedOrderServiceServer
	repo repository.OrderRepositoryI
}

func NewOrderControllerImpl(repo repository.OrderRepositoryI) OrderControllerI {
	return &OrderControllerImpl{repo: repo}
}

func (*OrderControllerImpl) CheckoutOrder(context.Context, *order.CheckoutOrderRequest) (*order.CheckoutOrderResponse, error) {
	panic("unimplemented")
}

func (*OrderControllerImpl) PayOrder(context.Context, *order.PayOrderRequest) (*emptypb.Empty, error) {
	panic("unimplemented")
}
