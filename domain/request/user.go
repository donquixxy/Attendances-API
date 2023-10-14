package request

import (
	"log"

	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// Read create user request given
func ReadRequest(c echo.Context) (*CreateUserRequest, error) {

	body := new(CreateUserRequest)

	if bindErr := c.Bind(body); bindErr != nil {
		log.Printf("failed to bind request user : %v", bindErr)
		return nil, bindErr
	}

	return body, nil
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

func ReadLoginRequest(c echo.Context) (*LoginRequest, error) {
	body := new(LoginRequest)

	if bindErr := c.Bind(body); bindErr != nil {
		log.Printf("failed to bind login user : %v", bindErr)
		return nil, bindErr
	}

	return body, nil
}
