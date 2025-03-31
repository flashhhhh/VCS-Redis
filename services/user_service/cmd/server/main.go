package main

import (
	"user_service/infrastructure/database"
	"user_service/infrastructure/grpc"
	"user_service/infrastructure/redis"
	"user_service/internal/handler"
	"user_service/internal/repository"
	"user_service/internal/service"

	"github.com/flashhhhh/pkg/env"
)

func main() {
	env.LoadEnv("config/user.env")

	dsn := "host=" + env.GetEnv("USER_DB_HOST", "localhost") + " user=" + env.GetEnv("USER_DB_USER", "root") + " password=" + env.GetEnv("USER_DB_PASSWORD", "") + " dbname=" + env.GetEnv("USER_DB_NAME", "user_service") + " port=" + env.GetEnv("USER_DB_PORT", "5432") + " sslmode=disable"
	db, err := database.ConnectDB(dsn)
	if err != nil {
		panic(err)
	}

	// Initialize Redis client
	redisAddr := env.GetEnv("REDIS_ADDR", "redis:6379")
	redisClient := redis.NewRedisClient(redisAddr)

	// Initialize the repository, service, and handler
	userRepo := repository.NewUserRepository(db, redisClient)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Start the gRPC server
	grpc.StartGRPCServer(userHandler)
}