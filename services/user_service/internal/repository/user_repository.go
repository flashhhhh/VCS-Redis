package repository

import "gorm.io/gorm"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: NewPostgresConnection(),
	}
}

func (r *UserRepository) CreateUser(username, password, name string) error {
	user := User{
		Username: username,
		Password: password,
		Name:     name,
	}

	// Create user in database
	result := r.db.Create(&user)
	return result.Error
}

func (r *UserRepository) Login(username string) (User, error) {
	user := User{}
	result := r.db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (r *UserRepository) GetUserByID(id int) (User, error) {
	user := User{}
	result := r.db.First(&user, id)

	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (User, error) {
	user := User{}
	result := r.db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (r *UserRepository) GetAllUsers() ([]User, error) {
	users := []User{}
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}