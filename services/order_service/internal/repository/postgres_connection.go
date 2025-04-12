package repository

import (
	"sync"

	"github.com/flashhhhh/pkg/env"
	"github.com/flashhhhh/pkg/gorm"
)

var db *gorm.DB
var lock = &sync.Mutex{}

func NewPostgresConnection() *gorm.DB {
	env.LoadEnv("config/order.env")
	
	if db == nil {
		lock.Lock()
		defer lock.Unlock()

		if db == nil {
			host := env.GetEnv("USER_DB_HOST", "localhost")
			port := env.GetEnv("USER_DB_PORT", "5432")
			user := env.GetEnv("USER_DB_USER", "postgres")
			password := env.GetEnv("USER_DB_PASSWORD", "password")
			dbname := env.GetEnv("USER_DB_NAME", "user")

			dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
			db = gorm.NewGormDB(dsn)
		}
	}

	return db
}