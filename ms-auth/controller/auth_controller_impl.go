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

func (ac *AuthControllerImpl) RegisterAuth(c context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	res, err := ac.repo.Insert(req)
	if err != nil {
		return &pb.RegisterResponse{}, err
	}
	return res, nil
}

func (ac *AuthControllerImpl) LoginAuth(c context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	res, err := ac.repo.FindUser(req)
	if err != nil {
		return &pb.LoginResponse{}, err
	}

	return res, nil
}
