package model

import (
	pb "ms-order/pb/order"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CheckoutResponse struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Products    []*pb.Product      `bson:"products"`
	PaymentId   int                `bson:"payment_id"`
	Payment     *pb.Payment        `bson:"payment"`
	VoucherId   int                `bson:"voucher_id"`
	Voucher     *pb.Voucher        `bson:"voucher"`
	TotalPrice  float64            `bson:"total_price"`
	OrderStatus string             `bson:"order_status"`
	OrderDate   string             `bson:"order_date"`
}

type CartProducts struct {
	ProductId     int     `json:"product_id"`
	Name          string  `json:"name"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	SubTotalPrice float64 `json:"sub_total_price"`
	Stock         int     `json:"stock"`
}

type Carts struct {
	ID            int     `json:"cart_id,omitempty" gorm:"column:cart_id"`
	UserId        int     `json:"user_id"`
	ProductId     int     `json:"product_id"`
	Quantity      int     `json:"quantity"`
	SubTotalPrice float64 `json:"sub_total_price"`
}

type Product struct {
	ID          int     `json:"product_id,omitempty" gorm:"column:product_id"`
	CategoryId  int     `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Stock       int64   `json:"stock"`
	Price       float64 `json:"price"`
}

type Payment struct {
	ID          int    `json:"payment_id,omitempty" gorm:"column:payment_id"`
	PaymentName string `json:"payment_name"`
}

type Voucher struct {
	ID          int    `json:"voucher_id,omitempty" gorm:"column:voucher_id"`
	VoucherName string `json:"voucher_name"`
}

type User struct {
	ID       int     `json:"user_id,omitempty" gorm:"column:user_id"`
	FullName string  `json:"full_name"`
	Email    string  `json:"email"`
	Balance  float64 `json:"balance"`
	Address  string  `json:"address"`
	Role     string  `json:"role,omitempty" gorm:"default:user"`
}
