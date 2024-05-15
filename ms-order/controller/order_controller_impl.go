package controller

import (
	"context"
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

func (oc *OrderControllerImpl) CheckoutOrder(c context.Context, pr *pb.CheckoutOrderRequest) (*pb.CheckoutOrderResponse, error) {
	res, err := oc.repo.CreateOrder(pr)
	if err != nil {
		return &pb.CheckoutOrderResponse{}, err
	}

	return res, nil
}

func (oc *OrderControllerImpl) PayOrder(c context.Context, pr *pb.PayOrderRequest) (*emptypb.Empty, error) {
	_, err := oc.repo.PayOrder(pr)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
