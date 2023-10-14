package server

import (
	"dg-test/ent"
	"dg-test/handler"
	"dg-test/repository"
	"dg-test/service"

	m "dg-test/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	E  *echo.Echo
	DB *ent.Client
}

func NewServer(
	db *ent.Client,
) *Server {
	e := echo.New()
	e.Use(middleware.CORS())

	// User repository

	s := &Server{
		E:  e,
		DB: db,
	}
	userRepository := repository.NewUserRepository(s.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	s.userRoute(userHandler)

	return s
}

// register user route to framework
func (s *Server) userRoute(handler handler.UserHandler) {
	group := s.E.Group("/api")

	group.POST("/user", handler.StoreUser, m.Auth("supersecret"))
	group.POST("/auth", handler.Login)
}

// func (s *Server) ErrorHandler (err error, c echo.Context) {
