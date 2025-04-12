package service

import (
	"errors"
	"time"
	"user_service/internal/repository"

	"github.com/flashhhhh/pkg/hash"
	"github.com/flashhhhh/pkg/jwt"
)

var UserRepository = repository.NewUserRepository()

func CreateUser(username, password, name string) error {
	return UserRepository.CreateUser(username, password, name)
}

func Login(username, password string) (string, error) {
	userInfo, err := UserRepository.Login(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !hash.CompareHashAndString(userInfo.Password, password) {
		return "", errors.New("invalid password")
	}

	token, _ := jwt.GenerateToken(map[string]any{
		"id": userInfo.ID,
	}, time.Minute * 5)

	return token, nil
}

func GetUserByID(id int) (map[string]any, error) {
	user, err := UserRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"id": user.ID,
		"username": user.Username,
		"name": user.Name,
	}, nil
}

func GetUserByUsername(username string) (map[string]any, error) {
	user, err := UserRepository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"id": user.ID,
		"username": user.Username,
		"name": user.Name,
	}, nil
}

func GetAllUsers() ([]map[string]any, error) {
	users, err := UserRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var userArray []map[string]any
	for _, user := range users {
		userArray = append(userArray, map[string]any{
			"id": user.ID,
			"username": user.Username,
			"name": user.Name,
		})
	}

	return userArray, nil
}