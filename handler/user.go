package handler

import (
	"context"
	"dg-test/domain/entity"
	"dg-test/domain/request"
	"dg-test/service"
	"dg-test/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	StoreUser(c echo.Context) error
	Login(c echo.Context) error
	UpdateUser(c echo.Context) error
	GetUsers(c echo.Context) error
}

type userHandler struct {
	s         service.UserService
	validator *validator.Validate
}

func NewUserHandler(
	s service.UserService,
	validator *validator.Validate,
) UserHandler {
	return &userHandler{
		s:         s,
		validator: validator,
	}
}

func (h *userHandler) GetUsers(c echo.Context) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	users, err := h.s.GetUsers(ctx)

	if err != nil {
		return ErrorResponse(err, c)
	}

	res := SuccessResponse("Successfully retrieved", users)

	return c.JSON(200, res)
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	body, err := request.ReadUserUpdateRequest(c)

	if err != nil {
		return ErrorResponse(err, c)
	}

	id := c.Param("id")

	result, err := h.s.UpdateUser(ctx, body, id)

	if err != nil {
		return ErrorResponse(err, c)
	}

	res := SuccessResponse("Success", result)

	return c.JSON(200, res)
}

func (h *userHandler) StoreUser(c echo.Context) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	body, err := request.ReadRequest(c)

	if err != nil {
		return ErrorResponse(err, c)
	}

	err = utils.ValidateStruct(body, h.validator)

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

	err = utils.ValidateStruct(body, h.validator)

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
