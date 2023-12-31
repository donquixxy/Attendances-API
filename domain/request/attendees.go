package request

import (
	"log"

	"github.com/labstack/echo/v4"
)

type CreateRequest struct {
	IDUser string `json:"id_user" form:"id_user" validate:"required"`
	Type   int    `json:"type" form:"type" validate:"required"`
}

func ReadCreateRequest(c echo.Context) (*CreateRequest, error) {
	body := new(CreateRequest)

	if bindErr := c.Bind(body); bindErr != nil {
		log.Printf("failed to bind request attendees %v", bindErr)
		return nil, bindErr
	}

	return body, nil
}

type UpdateAttendanceRequest struct {
	IDUser string `json:"id_user" form:"id_user"`
	Type   int    `json:"type" form:"type"`
}

func ReadUpdateAttendanceRequest(c echo.Context) (*UpdateAttendanceRequest, error) {
	body := new(UpdateAttendanceRequest)

	if bindErr := c.Bind(body); bindErr != nil {
		log.Printf("failed to bind update request %v", bindErr)
		return nil, bindErr
	}

	return body, nil
}
