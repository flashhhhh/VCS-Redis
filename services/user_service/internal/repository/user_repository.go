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
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
	redis *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepository {
	return &userRepository{db: db, redis: redis}
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}

	// Store user in Redis cache
	userJson, _ := json.Marshal(user)
	r.redis.Set(ctx, "user:"+strconv.Itoa(user.ID), userJson, 0)

	return nil
}

func (r *userRepository) Login(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User

	// Check if user is in Redis cache
	cachedUser, err := r.redis.Get(ctx, "user:"+username).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cachedUser), &user); err == nil {
			return &user, nil
		}
	}

	// If not in cache, query the database
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	// Store user in Redis cache
	userJson, _ := json.Marshal(user)
	r.redis.Set(ctx, "user:"+username, userJson, 0)
	r.redis.Set(ctx, "user:"+strconv.Itoa(user.ID), userJson, 0)

	return &user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User

	// Check if user is in Redis cache
	cachedUser, err := r.redis.Get(ctx, "user:"+strconv.Itoa(id)).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cachedUser), &user); err == nil {
			return &user, nil
		}
	}

	// If not in cache, query the database
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}

	// Store user in Redis cache
	userJson, _ := json.Marshal(user)
	r.redis.Set(ctx, "user:"+strconv.Itoa(user.ID), userJson, 0)
	
	return &user, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User

	// Check if user is in Redis cache
	cachedUser, err := r.redis.Get(ctx, "user:"+username).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cachedUser), &user); err == nil {
			return &user, nil
		}
	}

	// If not in cache, query the database
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	// Store user in Redis cache
	userJson, _ := json.Marshal(user)
	r.redis.Set(ctx, "user:"+username, userJson, 0)

	return &user, nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User

	// Check if users are in Redis cache
	cachedUsers, err := r.redis.Get(ctx, "users").Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cachedUsers), &users); err == nil {
			return users, nil
		}
	}

	// If not in cache, query the database
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	// Store users in Redis cache
	usersJson, _ := json.Marshal(users)
	r.redis.Set(ctx, "users", usersJson, 0)
	
	return users, nil
}