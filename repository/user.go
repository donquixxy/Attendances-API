package repository

import (
	"context"
	"dg-test/ent"
	"dg-test/ent/user"
	"log"
	"time"
)

// Contract for userrepository that needs to be satisfied
type UserRepository interface {
	CreateUser(ctx context.Context, value *ent.User) (*ent.User, error)
	FindUserByEmail(ctx context.Context, email string) (*ent.User, error)
	UpdateUser(ctx context.Context, value *ent.User) (*ent.User, error)
}

type userRepository struct {
	client *ent.Client
}

// User repository returns contract of userrepository
func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{
		client: client,
	}
}

func (s *userRepository) UpdateUser(ctx context.Context, value *ent.User) (*ent.User, error) {
	result := s.client.User.UpdateOneID(value.ID)

	// SetName(value.Name).
	// 	SetEmail(value.Email).SetPassword(value.Password).SetUpdatedAt(time.Now()).Save(ctx)

	if value.Name != "" {
		result.SetName(value.Name)
	}

	if value.Password != "" {
		result.SetPassword(value.Password)
	}

	if value.Email != "" {
		result.SetEmail(value.Email)
	}

	res, err := result.SetUpdatedAt(time.Now()).Save(ctx)

	return res, err
}

func (s *userRepository) CreateUser(ctx context.Context, value *ent.User) (*ent.User, error) {
	result, err := s.client.User.Create().SetID(value.ID).
		SetEmail(value.Email).SetName(value.Name).
		SetPassword(value.Password).
		SetCreatedAt(time.Now()).
		SetCreatedAt(time.Now()).Save(ctx)

	if err != nil {
		log.Printf("Failed create user : [%v]", err)
		return nil, err
	}

	return result, nil
}

func (s *userRepository) FindUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	result, err := s.client.User.Query().Where(user.Email(email)).Only(ctx)

	return result, err
}
