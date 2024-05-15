package controller

import (
	"context"
	pb "ms-user/pb/user"
	"ms-user/repository"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserControllerImpl struct {
	pb.UnimplementedUserServiceServer
	repo repository.UserRepositoryI
}

func NewUserControllerImpl(repo repository.UserRepositoryI) UserControllerI {
	return &UserControllerImpl{repo: repo}
}

func (uc *UserControllerImpl) AddProductToCart(c context.Context, req *pb.AddProductToCartRequest) (*emptypb.Empty, error) {
	_, err := uc.repo.AddProduct(req)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (uc *UserControllerImpl) GetCartItems(c context.Context, req *pb.GetCartItemsRequest) (*pb.GetCartItemsResponse, error) {
	res, err := uc.repo.FindCartItems(req)
	if err != nil {
		return &pb.GetCartItemsResponse{}, err
	}

	return res, nil
}

func (uc *UserControllerImpl) TopUp(c context.Context, req *pb.TopUpRequest) (*emptypb.Empty, error) {
	_, err := uc.repo.AddBalance(req)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
