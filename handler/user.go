package handler

import (
	"context"
	"dg-test/domain/entity"
	"dg-test/domain/request"
	"dg-test/service"
	"time"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	StoreUser(c echo.Context) error
}

type userHandler struct {
	s service.UserService
}

func NewUserHandler(
	s service.UserService,
) UserHandler {
	return &userHandler{
		s: s,
	}
}

func (h *userHandler) StoreUser(c echo.Context) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	body, err := request.ReadRequest(c)

	if err != nil {
		return c.JSON(400, ErrorResponse("failed to get request", err))
	}

	result, err := h.s.CreateUser(ctx, body)

	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, ErrorResponse("failed insert user", err))
	}

	res := SuccessResponse("Successfully created", result)

	return c.JSON(201, res)
}

func SuccessResponse(msg string, data any) *entity.Response {
	return &entity.Response{
		Msg:  msg,
		Data: data,
	}
}

func ErrorResponse(msg string, err error) *entity.ErrResponse {
	return &entity.ErrResponse{
		Msg:   msg,
		Error: err.Error(),
	}

}
