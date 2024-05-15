package repository

import (
	pb "ms-order/pb/order"

	"github.com/golang/protobuf/ptypes/empty"
)

type OrderRepositoryI interface {
	CreateOrder(*pb.CheckoutOrderRequest) (*pb.CheckoutOrderResponse, error)
	PayOrder(*pb.PayOrderRequest) (*empty.Empty, error)
}
