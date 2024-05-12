package controller

import (
	pb "ms-auth/pb/auth"

	"context"
)

type AuthControllerI interface {
	Register(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error)
}
