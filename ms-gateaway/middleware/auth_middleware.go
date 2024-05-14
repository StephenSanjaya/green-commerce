package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("authorization")

			if tokenString == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "token not found")
			}

			secret_key := os.Getenv("SECRET_KEY")
			parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid algo use")
				}
				return []byte(secret_key), nil
			})

			if parsedToken == nil || !parsedToken.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			if float64(time.Now().Unix()) > parsedToken.Claims.(jwt.MapClaims)["exp"].(float64) {
				return echo.NewHTTPError(http.StatusUnauthorized, "token expired: "+err.Error())
			}

			user_role := parsedToken.Claims.(jwt.MapClaims)["role"]
			if user_role == "user" && role == "admin" {
				return echo.NewHTTPError(http.StatusUnauthorized, "restricted endpoint: only admin")
			}

			c.Set("email", parsedToken.Claims.(jwt.MapClaims)["email"])
			c.Set("id", parsedToken.Claims.(jwt.MapClaims)["id"])

			return next(c)
		}
	}
}
