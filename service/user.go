package service

import (
	"context"
	"dg-test/domain/request"
	"dg-test/ent"
	"dg-test/exception"
	"dg-test/repository"
	"dg-test/token"
	"dg-test/utils"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService interface {
	CreateUser(ctx context.Context, r *request.CreateUserRequest) (*ent.User, error)
	Login(ctx context.Context, r *request.LoginRequest) (token, refreshToken string, err error)
	UpdateUser(ctx context.Context, r *request.UpdateRequest, idUser string) (*ent.User, error)
	GenerateToken(u *ent.User) (string, error)
	GenerateRefreshToken(u *ent.User) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) UpdateUser(ctx context.Context, r *request.UpdateRequest, idUser string) (*ent.User, error) {

	var hashedPassword string

	if r.Password != "" {
		r, err := utils.HashPassword(r.Password)

		if err != nil {
			return nil, &exception.BadRequestError{
				Message: err.Error(),
			}
		}

		hashedPassword = string(r)
	}

	u := &ent.User{
		ID:       idUser,
		Name:     r.Name,
		Email:    r.Email,
		Password: hashedPassword,
	}

	result, err := s.userRepository.UpdateUser(ctx, u)

	if err != nil {
		log.Printf("Failed update user [%v] ", err)
		return nil, err
	}

	return result, nil
}

func (s *userService) GenerateToken(u *ent.User) (string, error) {
	payload := &token.Payload{
		ID:    u.ID,
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "gusari",
			ExpiresAt: time.Now().AddDate(0, 0, 2).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := tokenClaims.SignedString([]byte("supersecret"))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GenerateRefreshToken(u *ent.User) (string, error) {
	payload := &token.Payload{
		ID:    u.ID,
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "gusari",
			ExpiresAt: time.Now().AddDate(0, 0, 12).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := tokenClaims.SignedString([]byte("supersecret2023012"))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) CreateUser(ctx context.Context, r *request.CreateUserRequest) (*ent.User, error) {

	hashedPassword, err := utils.HashPassword(r.Password)

	if err != nil {
		log.Printf("failed to hash password : %v", err)
		return nil, &exception.BadRequestError{
			Message: err.Error(),
		}
	}

	u := &ent.User{
		ID:        utils.GenerateUUID(),
		Name:      r.Name,
		Email:     r.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.userRepository.CreateUser(ctx, u)
}

func (s *userService) Login(ctx context.Context, r *request.LoginRequest) (token, refreshToken string, err error) {

	user, err := s.userRepository.FindUserByEmail(ctx, r.Email)

	if err != nil {
		return "", "", &exception.RecordNotFoundError{
			Message: err.Error(),
		}
	}

	// Validate password
	err = utils.ValidatePassword(r.Password, user.Password)

	if err != nil {
		return "", "", &exception.BadRequestError{
			Message: "wrong password",
		}
	}

	// Generate token
	t, err := s.GenerateToken(user)

	if err != nil {
		return "", "", &exception.BadRequestError{
			Message: err.Error(),
		}
	}

	// Refresh token
	re, err := s.GenerateRefreshToken(user)

	if err != nil {
		return "", "", &exception.BadRequestError{
			Message: err.Error(),
		}
	}

	return t, re, nil
}
