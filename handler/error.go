package handler

import (
	"dg-test/domain/entity"
	"dg-test/exception"

	"github.com/labstack/echo/v4"
)

const (
	badRequestStatusCode     = 400
	duplicateEntryStatusCode = 409
	notFoundStatusCode       = 404
	serviceErrorStatusCode   = 503
	succesStatusCode         = 200
	unauthorizedStatusCode   = 401
)

func ErrorResponse(err error, c echo.Context) error {

	res := &entity.ErrResponse{}

	if err != nil {
		switch err.(type) {
		case *exception.BadRequestError:
			{
				res.Error = err.Error()
				res.Msg = "Bad Request"
				return c.JSON(badRequestStatusCode, res)
			}
		case *exception.RecordNotFoundError:
			{
				res.Error = err.Error()
				res.Msg = "Not found"
				return c.JSON(notFoundStatusCode, res)
			}
		case *exception.UnauthorizedError:

			{
				res.Error = err.Error()
				res.Msg = "Unauthorized"
				return c.JSON(unauthorizedStatusCode, res)
			}
		default:
			{
				res.Error = err.Error()
				res.Msg = "Service Error"
				return c.JSON(serviceErrorStatusCode, res)
			}
		}
	}
	return nil
}
