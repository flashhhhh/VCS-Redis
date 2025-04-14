package service

import (
	"context"
	"errors"
	"time"
	"user_service/internal/domain"
	"user_service/internal/repository"

	"github.com/flashhhhh/pkg/hash"
	"github.com/flashhhhh/pkg/jwt"
)

type UserService interface {
	CreateUser(ctx context.Context, username, password, name string) error
	Login(ctx context.Context, username, password string) (string, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, username, password, name string) error {
	user := &domain.User{
		Username: username,
		Password: hash.HashString(password),
		Name:     name,
	}

	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.Login(ctx, username)
	if err != nil {
		return "", err
	}

	if !hash.CompareHashAndString(user.Password, password) {
		return "", errors.New("invalid password")
	}

	token, jwtErr := jwt.GenerateToken(map[string]any{
		"id": user.ID,
		"username": user.Username,
		"name": user.Name,
	}, time.Hour)

	return token, jwtErr
}

func (s *userService) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}