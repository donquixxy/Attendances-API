package middleware

import (
	"dg-test/domain/entity"
	"dg-test/token"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Auth(secret string) echo.MiddlewareFunc {

	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &token.Payload{},
		SigningKey:              []byte(secret),
		ErrorHandlerWithContext: JWTErrorHandler,
	})
}

func JWTErrorHandler(err error, c echo.Context) error {
	return c.JSON(401, entity.ErrResponse{
		Msg:   "Invalid token",
		Error: err.Error(),
	})
}

func JWTClaimsIdUser(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*token.Payload)
	return claims.ID
}

func JWTClaimsUser(c echo.Context) *token.Payload {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*token.Payload)
	return claims
}
