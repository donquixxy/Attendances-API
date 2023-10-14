package handler

import (
	"context"
	"dg-test/domain/request"
	"dg-test/service"
	"dg-test/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AttendeesHandler interface {
	InsertAttendees(c echo.Context) error
}

type attendeesHandler struct {
	service   service.AttendeesService
	validator *validator.Validate
}

func NewAttendeesHandler(service service.AttendeesService, validator *validator.Validate) AttendeesHandler {
	return &attendeesHandler{
		service:   service,
		validator: validator,
	}
}

func (h *attendeesHandler) InsertAttendees(c echo.Context) error {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	defer cancel()

	payload, err := request.ReadCreateRequest(c)

	if err != nil {
		return ErrorResponse(err, c)
	}

	err = utils.ValidateStruct(payload, h.validator)

	if err != nil {
		return ErrorResponse(err, c)
	}

	result, err := h.service.InsertAttendance(ctx, payload)

	if err != nil {
		return ErrorResponse(err, c)
	}

	res := SuccessResponse("Success", result)

	return c.JSON(201, res)
}
