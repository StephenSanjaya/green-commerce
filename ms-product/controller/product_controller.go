package controller

import (
	"context"
	pb "ms-product/pb/product"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductControllerI interface {
	GetAllProduct(context.Context, *emptypb.Empty) (*pb.ProductResponses, error)
	GetProduct(context.Context, *pb.ProductId) *pb.ProductResponse
	AddProduct(context.Context, *pb.ProductRequest) *pb.ProductResponse
	DeleteProduct(context.Context, *pb.ProductId) *emptypb.Empty
	UpdateProduct(context.Context, *pb.ProductId) *emptypb.Empty
}
