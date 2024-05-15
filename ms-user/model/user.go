package model

type User struct {
	ID       int     `json:"user_id,omitempty" gorm:"column:user_id"`
	FullName string  `json:"full_name"`
	Email    string  `json:"email"`
	Balance  float64 `json:"balance"`
	Address  string  `json:"address"`
	Role     string  `json:"role,omitempty" gorm:"default:user"`
}
