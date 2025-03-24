package repository

import (
	"github.com/flashhhhh/pkg/hash"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name    string `json:"name"`
}

func CreateUser(userInfo map[string]any) (User, error) {
	username := userInfo["username"].(string)
	password := userInfo["password"].(string)
	name := userInfo["name"].(string)

	user := User{
		Username: username,
		Password: hash.HashString(password),
		Name:    name,
	}

	// Create user in database
	db := NewPostgresConnection()
	result := db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}