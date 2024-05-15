package model

type Cart struct {
	ID            int     `json:"cart_id,omitempty" gorm:"column:cart_id"`
	ProductId     int     `json:"product_id"`
	UserId        int     `json:"user_id"`
	Quantity      int     `json:"quantity"`
	SubTotalPrice float64 `json:"sub_total_price"`
}
