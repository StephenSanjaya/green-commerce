package repository

import (
	"errors"
	"ms-product/model"
	pb "ms-product/pb/product"
	"net/http"

	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) ProductRepositoryI {
	return &ProductRepositoryImpl{db: db}
}

func (pr *ProductRepositoryImpl) Create(req *pb.ProductRequest) (*pb.ProductResponse, error) {
	product := &model.Product{
		CategoryId:  int(req.CategoryId),
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
	}
	res := pr.db.Create(product)
	if res.Error != nil {
		return &pb.ProductResponse{}, status.Errorf(http.StatusInternalServerError, res.Error.Error())
	}

	return &pb.ProductResponse{
		ProductId:   int64(product.ID),
		CategoryId:  int64(product.CategoryId),
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
	}, nil
}

func (pr *ProductRepositoryImpl) Delete(id int) error {
	res := pr.db.Delete(&model.Product{}, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return status.Errorf(http.StatusNotFound, res.Error.Error())
	}
	if res.Error != nil {
		return status.Errorf(http.StatusInternalServerError, res.Error.Error())
	}

	return nil
}

func (pr *ProductRepositoryImpl) FindAll() (*pb.ProductResponses, error) {
	products := new([]model.Product)

	res := pr.db.Find(products)
	if res.Error != nil {
		return &pb.ProductResponses{}, status.Errorf(http.StatusInternalServerError, res.Error.Error())
	}

	y := pb.ProductResponses{}
	for _, p := range *products {
		x := pb.ProductResponse{
			ProductId:   int64(p.ID),
			CategoryId:  int64(p.CategoryId),
			Name:        p.Name,
			Description: p.Description,
			Stock:       p.Stock,
			Price:       p.Price,
		}
		y.Products = append(y.Products, &x)
	}

	return &y, nil
}

func (pr *ProductRepositoryImpl) FindOne(id int) (*pb.ProductResponse, error) {
	product := new(model.Product)
	res := pr.db.Find(product, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &pb.ProductResponse{}, status.Errorf(http.StatusNotFound, res.Error.Error())
	}
	if res.Error != nil {
		return &pb.ProductResponse{}, status.Errorf(http.StatusInternalServerError, res.Error.Error())
	}

	return &pb.ProductResponse{
		ProductId:   int64(product.ID),
		CategoryId:  int64(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
	}, nil
}

func (pr *ProductRepositoryImpl) Update(id int, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	product := new(model.Product)
	res := pr.db.Model(product).Where("product_id = ?", id).Updates(model.Product{
		CategoryId:  int(req.CategoryId),
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
	})
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &pb.ProductResponse{}, status.Errorf(http.StatusNotFound, res.Error.Error())
	}
	if res.Error != nil {
		return &pb.ProductResponse{}, status.Errorf(http.StatusInternalServerError, res.Error.Error())
	}

	return &pb.ProductResponse{
		ProductId:   int64(product.ID),
		CategoryId:  int64(product.CategoryId),
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
	}, nil
}
