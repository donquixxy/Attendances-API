package handler

import (
	"context"
	"dg-test/domain/request"
	"dg-test/middleware"
	"dg-test/service"
	"dg-test/utils"
	"log"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AttendeesHandler interface {
	InsertAttendees(c echo.Context) error // Id user input
	StoreAttendance(c echo.Context) error // Id user claims by jwt
	GetAllAttendances(c echo.Context) error
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

func (h *attendeesHandler) GetAllAttendances(c echo.Context) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	defer cancel()

	result, err := h.service.GetAllAttendances(ctx)

	if err != nil {
		return ErrorResponse(err, c)
	}

	res := SuccessResponse("Success", result)

	return c.JSON(200, res)
}

func (h *attendeesHandler) StoreAttendance(c echo.Context) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	defer cancel()

	user := middleware.JWTClaimsUser(c)

	typ := c.FormValue("type")

	t, err := strconv.Atoi(typ)

	if err != nil {
		log.Printf("Failed to convert atoi %v ", err)
		return ErrorResponse(err, c)
	}

	result, err := h.service.StoreAttendance(ctx, user.ID, t)

	if err != nil {
		return ErrorResponse(err, c)
	}

	res := SuccessResponse("Succesfullly absent", result)

	return c.JSON(201, res)
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
