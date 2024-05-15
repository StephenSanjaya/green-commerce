package controller

import (
	pb "ms-gateaway/pb/order"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
)

type OrderControllerImpl struct {
	orderGRPC pb.OrderServiceClient
}

func NewOrderControllerImpl(orderGRPC pb.OrderServiceClient) OrderControllerI {
	return &OrderControllerImpl{orderGRPC: orderGRPC}
}

func (oc *OrderControllerImpl) CheckoutOrder(c echo.Context) error {
	req := &pb.CheckoutOrderRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request: "+err.Error())
	}

	res, err := oc.orderGRPC.CheckoutOrder(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "failed to checkout order: "+err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success checkout order",
		"order":   res,
	})
}

func (oc *OrderControllerImpl) PayOrder(c echo.Context) error {
	user_id := int(c.Get("id").(float64))
	order_id := c.Param("order_id")
	req := &pb.PayOrderRequest{
		UserId:  int64(user_id),
		OrderId: order_id,
	}

	_, err := oc.orderGRPC.PayOrder(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(int(status.Code(err)), "failed to pay order: "+err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success pay order",
	})
}
