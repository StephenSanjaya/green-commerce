package repository

import (
	"context"
	"errors"
	"fmt"
	"ms-order/model"
	pb "ms-order/pb/order"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	mongo *mongo.Collection
	pg    *gorm.DB
}

func NewOrderRepositoryImpl(mongo *mongo.Collection, pg *gorm.DB) OrderRepositoryI {
	return &OrderRepositoryImpl{mongo: mongo, pg: pg}
}

func (or *OrderRepositoryImpl) CreateOrder(req *pb.CheckoutOrderRequest) (*pb.CheckoutOrderResponse, error) {
	products := new([]model.CartProducts)
	res := or.pg.Model(&model.Carts{}).Select("products.product_id, products.name, carts.quantity, products.price, carts.sub_total_price").Joins("join products on products.product_id = carts.product_id").Scan(products)
	if res.RowsAffected == 0 {
		return &pb.CheckoutOrderResponse{}, status.Errorf(http.StatusBadRequest, "no item in your cart")
	}
	if res.Error != nil {
		return &pb.CheckoutOrderResponse{}, status.Errorf(http.StatusInternalServerError, res.Error.Error())
	}

	totalPrice := 0.0
	pb_product := []*pb.Product{}

	for _, p := range *products {
		l := &pb.Product{
			ProductId:     int64(p.ProductId),
			ProductName:   p.Name,
			Quantity:      int64(p.Quantity),
			Price:         p.Price,
			SubTotalPrice: p.SubTotalPrice,
		}
		totalPrice += p.SubTotalPrice
		pb_product = append(pb_product, l)
	}

	payment := new(model.Payment)
	resPayment := or.pg.First(payment, req.PaymentId)
	if errors.Is(resPayment.Error, gorm.ErrRecordNotFound) {
		return &pb.CheckoutOrderResponse{}, status.Errorf(http.StatusNotFound, "payment id not found")
	}
	if resPayment.Error != nil {
		return &pb.CheckoutOrderResponse{}, status.Errorf(http.StatusInternalServerError, resPayment.Error.Error())
	}

	voucher := new(model.Voucher)
	resVoucher := or.pg.First(voucher, req.VoucherId)
	if errors.Is(resVoucher.Error, gorm.ErrRecordNotFound) {
		return &pb.CheckoutOrderResponse{}, status.Errorf(http.StatusNotFound, "voucher id not found")
	}
	if resVoucher.Error != nil {
		return &pb.CheckoutOrderResponse{}, status.Errorf(http.StatusInternalServerError, resVoucher.Error.Error())
	}
	if req.VoucherId == 1 {
		totalPrice = totalPrice - (totalPrice * 0.1)
	} else if req.VoucherId == 2 {
		totalPrice = totalPrice - (totalPrice * 0.2)
	}

	pb_payment := &pb.Payment{}
	pb_payment.PaymentName = payment.PaymentName

	pb_voucher := &pb.Voucher{}
	pb_voucher.VoucherName = voucher.VoucherName

	x := &pb.CheckoutOrderResponse{
		Products:    pb_product,
		PaymentId:   req.PaymentId,
		Payment:     pb_payment,
		VoucherId:   req.VoucherId,
		Voucher:     pb_voucher,
		TotalPrice:  totalPrice,
		OrderStatus: "pending",
		OrderDate:   time.Now().Format("2006-01-02"),
	}
	lk := &model.CheckoutResponse{
		Products:    pb_product,
		PaymentId:   int(req.PaymentId),
		Payment:     pb_payment,
		VoucherId:   int(req.VoucherId),
		Voucher:     pb_voucher,
		TotalPrice:  totalPrice,
		OrderStatus: "pending",
		OrderDate:   time.Now().Format("2006-01-02"),
	}

	resInsert, err := or.mongo.InsertOne(context.TODO(), lk)
	if err != nil {
		return &pb.CheckoutOrderResponse{}, status.Errorf(http.StatusInternalServerError, err.Error())
	}
	x.OrderId = resInsert.InsertedID.(primitive.ObjectID).Hex()

	resDel := or.pg.Where("user_id = ?", req.UserId).Delete(&model.Carts{})
	if resDel.Error != nil {
		return &pb.CheckoutOrderResponse{}, status.Errorf(http.StatusInternalServerError, resDel.Error.Error())
	}

	return x, nil
}

func (or *OrderRepositoryImpl) PayOrder(req *pb.PayOrderRequest) (*emptypb.Empty, error) {
	user := new(model.User)
	resUser := or.pg.First(user, req.UserId)
	if errors.Is(resUser.Error, gorm.ErrRecordNotFound) {
		return &emptypb.Empty{}, status.Errorf(http.StatusNotFound, "user id not found")
	}
	if resUser.Error != nil {
		return &emptypb.Empty{}, status.Errorf(http.StatusInternalServerError, resUser.Error.Error())
	}

	res := or.mongo.FindOne(context.Background(), bson.M{"order_id": req.OrderId})
	if res.Err() != nil {
		return &emptypb.Empty{}, status.Errorf(http.StatusInternalServerError, "failed to get order")
	}

	fmt.Println("res: ", res)

	// JSONData := struct {
	// 	TotalPrice float64 `bson:"total_price"`
	// }{}

	// decodeError := res.Decode(&JSONData)
	// if decodeError != nil {
	// 	return &emptypb.Empty{}, status.Errorf(http.StatusInternalServerError, "failed to decode")
	// }

	// if user.Balance < JSONData.TotalPrice {
	// 	return &emptypb.Empty{}, status.Errorf(http.StatusInternalServerError, "balance not enough")
	// }

	return &emptypb.Empty{}, nil
}
