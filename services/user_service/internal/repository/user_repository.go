package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"user_service/internal/domain"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, username string) (*domain.User, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
	redis *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepository {
	return &userRepository{
		db:    db,
		redis: redis,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	err := r.db.Create(user).Error

	if err != nil {
		return err
	}

	// Cache the user in Redis
	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = r.redis.Set(ctx, "user:"+strconv.Itoa(user.ID), userData, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Login(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User

	// Check Redis cache first
	cachedUser, err := r.redis.Get(ctx, "user:"+strconv.Itoa(id)).Result()
	if err == nil {
		// User found in cache, unmarshal it
		err = json.Unmarshal([]byte(cachedUser), &user)
		if err == nil {
			return &user, nil
		}
	}

	// If not found in cache, query the database
	err = r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	// Cache the user in Redis
	userData, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	err = r.redis.Set(ctx, "user:"+strconv.Itoa(id), userData, 0).Err()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}