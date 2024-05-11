package helper

import (
	pb "ms-gateaway/pb/auth"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWT(user *pb.LoginResponse) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret_key := os.Getenv("SECRET_KEY")
	secret_token := []byte(secret_key)

	tokenString, err := token.SignedString(secret_token)

	return tokenString, err
}
