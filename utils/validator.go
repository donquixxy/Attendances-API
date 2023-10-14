package utils

import (
	"dg-test/exception"

	"github.com/go-playground/validator/v10"
)

// Validate any value of v
func ValidateStruct(v any, validator *validator.Validate) error {
	err := validator.Struct(v)

	if err != nil {
		return &exception.BadRequestError{
			Message: err.Error(),
		}
	}

	return nil
}
