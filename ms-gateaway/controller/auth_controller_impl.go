package controller

import (
    "ms-gateaway/helper"
    pb "ms-gateaway/pb/auth"
    "net/http"

    "github.com/labstack/echo/v4"
    "google.golang.org/grpc/status"
)

type AuthControllerImpl struct {
    authGRPC pb.AuthServiceClient
}

func NewAuthController(authGRPC pb.AuthServiceClient) AuthControllerI {
    return &AuthControllerImpl{authGRPC: authGRPC}
}

// RegisterAuth godoc
// @Summary Register a new user
// @Description Register a new user with full details
// @Tags auth
// @Accept json
// @Produce json
// @Param registerRequest body pb.RegisterRequest true "Register Request"
// @Success 201 {object} pb.RegisterResponse
// @Failure 400 {object} helper.HTTPError
// @Failure 500 {object} helper.HTTPError
// @Router /register [post]
func (ac *AuthControllerImpl) RegisterAuth(c echo.Context) error {
    req := &pb.RegisterRequest{}
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid body request: "+err.Error())
    }

    hashedPass := helper.HashedPassword(req.Password)
    req.Password = hashedPass

    res, err := ac.authGRPC.RegisterAuth(c.Request().Context(), req)
    if err != nil {
        return echo.NewHTTPError(int(status.Code(err)), "invalid body request: "+err.Error())
    }

    return c.JSON(http.StatusCreated, echo.Map{
        "message": "success register",
        "user":    res,
    })
}

// LoginAuth godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body pb.LoginRequest true "Login Request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} helper.HTTPError
// @Failure 500 {object} helper.HTTPError
// @Router /login [post]
func (ac *AuthControllerImpl) LoginAuth(c echo.Context) error {
    req := &pb.LoginRequest{}
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid body request: "+err.Error())
    }

    res, err := ac.authGRPC.LoginAuth(c.Request().Context(), req)
    if err != nil {
        return echo.NewHTTPError(int(status.Code(err)), "login failed: "+err.Error())
    }

    isPassword := helper.CompareHashedPassword(req.Password, res.Password)
    if !isPassword {
        return echo.NewHTTPError(http.StatusBadRequest, "password not match")
    }

    token, err := helper.CreateJWT(res)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to create token: "+err.Error())
    }

    return c.JSON(http.StatusOK, echo.Map{
        "message": "success login",
        "jwt":     token,
    })
}
