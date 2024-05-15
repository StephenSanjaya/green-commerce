package model

type Product struct {
	ID          int     `json:"product_id,omitempty" gorm:"column:product_id"`
	CategoryId  int     `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Stock       int64   `json:"stock"`
	Price       float64 `json:"price"`
}