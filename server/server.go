package server

import (
	"dg-test/domain/entity"
	"dg-test/ent"
	"dg-test/handler"
	"dg-test/repository"
	"dg-test/service"

	m "dg-test/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	E         *echo.Echo
	DB        *ent.Client
	secretKey string
	Validator *validator.Validate
}

func NewServer(
	db *ent.Client,
) *Server {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = ErrorHandler
	// User repository

	s := &Server{
		E:         e,
		DB:        db,
		secretKey: "supersecret",
		Validator: validator.New(),
	}
	userRepository := repository.NewUserRepository(s.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, s.Validator)

	attendeesRepository := repository.NewAttendeesRepository(s.DB)
	attendeesService := service.NewAttendeesService(attendeesRepository)
	attendeesHandler := handler.NewAttendeesHandler(attendeesService, s.Validator)

	s.userRoute(userHandler)
	s.attendeesRoute(attendeesHandler)

	return s
}

// register user route to framework
func (s *Server) userRoute(handler handler.UserHandler) {
	group := s.E.Group("/api")

	group.POST("/user", handler.StoreUser, m.Auth(s.secretKey))
	group.POST("/auth", handler.Login)
	group.PUT("/user/:id", handler.UpdateUser, m.Auth(s.secretKey))
	group.DELETE("/user/:id", handler.DeleteUser, m.Auth(s.secretKey))
	group.GET("/user/:id", handler.GetUserByID, m.Auth(s.secretKey))
	group.GET("/users", handler.GetUsers, m.Auth(s.secretKey))
}

func (s *Server) attendeesRoute(handler handler.AttendeesHandler) {
	group := s.E.Group("/api", m.Auth(s.secretKey))

	group.POST("/attendance", handler.InsertAttendees)
}

func ErrorHandler(err error, c echo.Context) {
	js := &entity.ErrResponse{}

	js.Error = err.Error()
	js.Msg = "Internal Server error"
	c.JSON(500, js)
}
