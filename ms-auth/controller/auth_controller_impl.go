package controller

import (
	"context"
	pb "ms-auth/pb/auth"
	"ms-auth/repository"
)

type AuthControllerImpl struct {
	pb.UnimplementedAuthServiceServer
	repo repository.AuthRepositoryI
}

func NewAuthControllerImpl(repo repository.AuthRepositoryI) AuthControllerI {
	return &AuthControllerImpl{repo: repo}
}

func (ac *AuthControllerImpl) Register(c context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	res, err := ac.repo.Insert(req)
	if err != nil {
		return &pb.RegisterResponse{}, err
	}

	return &pb.RegisterResponse{
		UserId:   res.UserId,
		FullName: res.FullName,
		Email:    res.Email,
		Balance:  res.Balance,
		Address:  res.Address,
		Role:     res.Role,
	}, nil
}

func (ac *AuthControllerImpl) Login(c context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	res, err := ac.repo.FindUser(req)
	if err != nil {
		return &pb.LoginResponse{}, err
	}

	return &pb.LoginResponse{
		UserId:   res.UserId,
		Email:    res.Email,
		Password: res.Password,
	}, nil
}
