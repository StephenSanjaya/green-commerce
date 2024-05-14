package controller

import (
	"context"
	pb "ms-product/pb/product"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductControllerI interface {
	GetAllProduct(context.Context, *emptypb.Empty) (*pb.ProductResponses, error)
	GetProduct(context.Context, *pb.ProductId) (*pb.ProductResponse, error)
	AddProduct(context.Context, *pb.ProductRequest) (*pb.ProductResponse, error)
	UpdateProduct(context.Context, *pb.ProductRequest) (*pb.ProductResponse, error)
	DeleteProduct(context.Context, *pb.ProductId) (*emptypb.Empty, error)
}
