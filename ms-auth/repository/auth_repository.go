package repository

import pb "ms-auth/pb/auth"

type AuthRepositoryI interface {
	Insert(*pb.RegisterRequest) (*pb.RegisterResponse, error)
	FindUser(*pb.LoginRequest) (*pb.LoginResponse, error)
}
