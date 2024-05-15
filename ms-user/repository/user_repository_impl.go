package repository

import (
	"errors"
	"ms-user/model"
	pb "ms-user/pb/user"
	"net/http"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepositoryI {
	return &UserRepositoryImpl{db: db}
}

func (uc *UserRepositoryImpl) AddBalance(req *pb.TopUpRequest) (*emptypb.Empty, error) {
	user := new(model.User)
	res := uc.db.First(user, req.UserId)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &emptypb.Empty{}, status.Errorf(http.StatusNotFound, res.Error.Error())
	}
	if res.Error != nil {
		return &emptypb.Empty{}, status.Errorf(http.StatusInternalServerError, res.Error.Error())
	}

	updatedBalance := user.Balance + req.Amount

	resUpdate := uc.db.Model(&model.User{}).Where("user_id = ?", req.UserId).Update("balance", updatedBalance)
	if res.Error != nil {
		return &emptypb.Empty{}, status.Errorf(http.StatusInternalServerError, resUpdate.Error.Error())
	}

	return &emptypb.Empty{}, nil
}

func (uc *UserRepositoryImpl) AddProduct(req *pb.AddProductToCartRequest) (*emptypb.Empty, error) {
	product := new(model.Product)
	res := uc.db.First(product, req.ProductId)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &emptypb.Empty{}, status.Errorf(http.StatusNotFound, res.Error.Error())
	}
	if res.Error != nil {
		return &emptypb.Empty{}, status.Errorf(http.StatusInternalServerError, res.Error.Error())
	}

	req.SubTotalPrice = float64(req.Quantity) * product.Price

	cartReq := &model.Cart{
		ProductId:     int(req.ProductId),
		UserId:        int(req.UserId),
		Quantity:      int(req.Quantity),
		SubTotalPrice: req.SubTotalPrice,
	}

	resCreate := uc.db.Create(cartReq)
	if resCreate.Error != nil {
		return &emptypb.Empty{}, status.Errorf(http.StatusInternalServerError, resCreate.Error.Error())
	}

	return &emptypb.Empty{}, nil
}

func (uc *UserRepositoryImpl) FindCartItems(req *pb.GetCartItemsRequest) (*pb.GetCartItemsResponse, error) {
	cartItems := new([]model.Cart)

	res := uc.db.Where("user_id = ?", req.UserId).Find(cartItems)
	if res != nil {
		return &pb.GetCartItemsResponse{}, status.Errorf(http.StatusInternalServerError, res.Error.Error())
	}

	y := pb.GetCartItemsResponse{}
	for _, i := range *cartItems {
		x := pb.Cart{
			ProductId:     int64(i.ID),
			Quantity:      int64(i.Quantity),
			SubTotalPrice: i.SubTotalPrice,
		}
		y.Carts = append(y.Carts, &x)
	}

	return &y, nil
}
