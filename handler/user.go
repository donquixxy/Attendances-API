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
	Login(c echo.Context) error
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
		return ErrorResponse(err, c)
	}

	result, err := h.s.CreateUser(ctx, body)

	if err != nil {
		return ErrorResponse(err, c)
	}

	res := SuccessResponse("Successfully created", result)

	return c.JSON(201, res)
}

func (h *userHandler) Login(c echo.Context) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	body, err := request.ReadLoginRequest(c)

	if err != nil {
		return ErrorResponse(err, c)
	}

	token, refreshToken, err := h.s.Login(ctx, body)

	if err != nil {
		return ErrorResponse(err, c)
	}

	return c.JSON(200, SuccessResponse("Success", map[string]interface{}{
		"access_token":  token,
		"refresh_token": refreshToken,
	}))

}

func SuccessResponse(msg string, data any) *entity.Response {
	return &entity.Response{
		Msg:  msg,
		Data: data,
	}
}
