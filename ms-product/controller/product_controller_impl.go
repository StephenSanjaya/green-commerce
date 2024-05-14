package controller

import (
	"context"
	pb "ms-product/pb/product"
	"ms-product/repository"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductControllerImpl struct {
	pb.UnimplementedProductServiceServer
	repo repository.ProductRepositoryI
}

func NewProductControllerImpl(repo repository.ProductRepositoryI) ProductControllerI {
	return &ProductControllerImpl{repo: repo}
}

func (pc *ProductControllerImpl) GetAllProduct(c context.Context, empty *emptypb.Empty) (*pb.ProductResponses, error) {
	res, err := pc.repo.FindAll()
	if err != nil {
		return &pb.ProductResponses{}, err
	}
	return res, nil
}

func (pc *ProductControllerImpl) GetProduct(c context.Context, req_id *pb.ProductId) (*pb.ProductResponse, error) {
	res, err := pc.repo.FindOne(int(req_id.ProductId))
	if err != nil {
		return &pb.ProductResponse{}, err
	}
	return res, nil
}

func (pc *ProductControllerImpl) AddProduct(c context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	res, err := pc.repo.Create(req)
	if err != nil {
		return &pb.ProductResponse{}, err
	}
	return res, nil
}

func (pc *ProductControllerImpl) UpdateProduct(c context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	product_id := c.Value("product_id")
	res, err := pc.repo.Update(product_id.(int), req)
	if err != nil {
		return &pb.ProductResponse{}, err
	}
	return res, nil
}

func (pc *ProductControllerImpl) DeleteProduct(c context.Context, req_id *pb.ProductId) (*emptypb.Empty, error) {
	err := pc.repo.Delete(int(req_id.ProductId))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
