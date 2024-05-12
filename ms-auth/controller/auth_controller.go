package controller

import (
	pb "ms-auth/pb/auth"

	"context"
)

type AuthControllerI interface {
	RegisterAuth(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error)
	LoginAuth(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error)
}
