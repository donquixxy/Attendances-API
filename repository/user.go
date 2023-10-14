package repository

import (
	"context"
	"dg-test/ent"
	"log"
	"time"
)

// Contract for userrepository that needs to be satisfied
type UserRepository interface {
	CreateUser(ctx context.Context, value *ent.User) (*ent.User, error)
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
