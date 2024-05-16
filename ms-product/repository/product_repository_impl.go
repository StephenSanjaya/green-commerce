package repository

import (
	"context"
	"encoding/json"
	"ms-product/model"
	pb "ms-product/pb/product"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewProductRepositoryImpl(db *gorm.DB, redisClient *redis.Client) ProductRepositoryI {
	return &ProductRepositoryImpl{db: db, redis: redisClient}
}

func (pr *ProductRepositoryImpl) FindAll() (*pb.ProductResponses, error) {
	ctx := context.Background()
	cacheKey := "products:all"

	cachedData, err := pr.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedProducts pb.ProductResponses
		if err := json.Unmarshal([]byte(cachedData), &cachedProducts); err == nil {
			return &cachedProducts, nil
		}
	}

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

	dataToCache, err := json.Marshal(&y)
	if err == nil {
		pr.redis.Set(ctx, cacheKey, dataToCache, time.Hour).Err()
	}

	return &y, nil
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

func (pr *ProductRepositoryImpl) FindOne(id int) (*pb.ProductResponse, error) {
	product := new(model.Product)
	res := pr.db.Find(product, id)
	if res.RowsAffected == 0 {
		return &pb.ProductResponse{}, status.Errorf(http.StatusNotFound, "product id not found")
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
	if res.RowsAffected == 0 {
		return &pb.ProductResponse{}, status.Errorf(http.StatusNotFound, "product id not found")
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

func (pr *ProductRepositoryImpl) Delete(id int) error {
	res := pr.db.Delete(&model.Product{}, id)
	if res.RowsAffected == 0 {
		return status.Errorf(http.StatusNotFound, "product id not found")
	}
	if res.Error != nil {
		return status.Errorf(http.StatusInternalServerError, res.Error.Error())
	}

	return nil
}
