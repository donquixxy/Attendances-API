package service

import (
	"context"
	"dg-test/domain/request"
	"dg-test/ent"
	"dg-test/exception"
	"dg-test/repository"
	"dg-test/utils"
	"log"
	"time"
)

type UserService interface {
	CreateUser(ctx context.Context, r *request.CreateUserRequest) (*ent.User, error)
	Login(ctx context.Context, r *request.LoginRequest) (*ent.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
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

func (s *userService) Login(ctx context.Context, r *request.LoginRequest) (*ent.User, error) {

	user, err := s.userRepository.FindUserByEmail(ctx, r.Email)

	if err != nil {
		return nil, &exception.RecordNotFoundError{
			Message: err.Error(),
		}
	}

	// Validate password
	err = utils.ValidatePassword(r.Password, user.Password)

	if err != nil {
		return nil, &exception.BadRequestError{
			Message: "wrong password",
		}
	}

	return user, nil
}
