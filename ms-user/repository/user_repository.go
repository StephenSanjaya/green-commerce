package repository

import (
	pb "ms-user/pb/user"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserRepositoryI interface {
	AddProduct(*pb.AddProductToCartRequest) (*emptypb.Empty, error)
	AddBalance(*pb.TopUpRequest) (*emptypb.Empty, error)
	FindCartItems(*pb.GetCartItemsRequest) (*pb.GetCartItemsResponse, error)
}
