package service

import (
	"user_service/internal/repository"
)

func CreateUser(userInfo map[string]any) error {
	_, err := repository.CreateUser(userInfo)
	return err;
}