package repository

import (
	"errors"
	"ms-auth/model"
	pb "ms-auth/pb/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) AuthRepositoryI {
	return &AuthRepositoryImpl{db: db}
}

func (ar *AuthRepositoryImpl) Insert(req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	userReq := &model.User{
		FullName: req.FullName,
		Email:    req.Email,
		Balance:  0,
		Address:  req.Address,
		Role:     req.Role,
	}
	res := ar.db.Create(userReq)
	if res.Error != nil {
		return &pb.RegisterResponse{}, status.Errorf(codes.Internal, res.Error.Error())
	}

	return &pb.RegisterResponse{
		UserId:   userReq.UserId,
		FullName: userReq.FullName,
		Email:    userReq.Email,
		Balance:  userReq.Balance,
		Address:  userReq.Address,
		Role:     userReq.Role,
	}, nil
}

func (ar *AuthRepositoryImpl) FindUser(req *pb.LoginRequest) (*pb.LoginResponse, error) {
	loginRes := new(pb.LoginResponse)
	res := ar.db.Where("email = ?", req.Email).First(loginRes)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &pb.LoginResponse{}, status.Errorf(codes.NotFound, res.Error.Error())
	}
	if res.Error != nil {
		return &pb.LoginResponse{}, status.Errorf(codes.Internal, res.Error.Error())
	}

	return loginRes, nil
}
