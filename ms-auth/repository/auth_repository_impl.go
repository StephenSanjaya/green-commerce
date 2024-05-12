package repository

import (
	"errors"
	"ms-auth/model"
	pb "ms-auth/pb/auth"
	"net/http"

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
		Password: req.Password,
		Address:  req.Address,
	}
	res := ar.db.Create(userReq)
	if res.Error != nil {
		return &pb.RegisterResponse{}, status.Errorf(http.StatusInternalServerError, res.Error.Error())
	}

	return &pb.RegisterResponse{
		UserId:   int64(userReq.ID),
		FullName: userReq.FullName,
		Email:    userReq.Email,
		Balance:  userReq.Balance,
		Address:  userReq.Address,
		Role:     userReq.Role,
	}, nil
}

func (ar *AuthRepositoryImpl) FindUser(req *pb.LoginRequest) (*pb.LoginResponse, error) {
	loginRes := new(pb.LoginResponse)
	res := ar.db.Model(&model.User{}).Where("email = ?", req.Email).First(loginRes)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &pb.LoginResponse{}, status.Errorf(http.StatusNotFound, res.Error.Error())
	}
	if res.Error != nil {
		return &pb.LoginResponse{}, status.Errorf(http.StatusInternalServerError, res.Error.Error())
	}

	return loginRes, nil
}
