package model

type User struct {
	UserId   int64   `json:"user_id,omitempty"`
	FullName string  `json:"full_name,omitempty"`
	Email    string  `json:"email,omitempty"`
	Balance  float32 `json:"balance,omitempty"`
	Address  string  `json:"address,omitempty"`
	Role     string  `json:"role,omitempty"`
}
