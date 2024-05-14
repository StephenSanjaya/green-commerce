package repository

import pb "ms-product/pb/product"

type ProductRepositoryI interface {
	FindAll() (*pb.ProductResponses, error)
	FindOne(id int) (*pb.ProductResponse, error)
	Create(*pb.ProductRequest) (*pb.ProductResponse, error)
	Update(int, *pb.ProductRequest) (*pb.ProductResponse, error)
	Delete(id int) error
}
