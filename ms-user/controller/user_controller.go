package controller

import (
	"context"
	pb "ms-user/pb/user"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserControllerI interface {
	AddProductToCart(context.Context, *pb.AddProductToCartRequest) (*emptypb.Empty, error)
	TopUp(context.Context, *pb.TopUpRequest) (*emptypb.Empty, error)
	GetCartItems(context.Context, *pb.GetCartItemsRequest) (*pb.GetCartItemsResponse, error)
}
