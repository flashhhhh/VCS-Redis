package main

import (
	"log"
	"os"
	"user_service/infrastructure/grpc"
	database "user_service/infrastructure/postgres"
	"user_service/infrastructure/redis"
	"user_service/internal/handler"
	"user_service/internal/repository"
	"user_service/internal/service"

	"github.com/flashhhhh/pkg/env"
)

func main() {
	var environment string

	args := os.Args
	if len(args) > 1 {
		environment = args[1]
		println("Running in environment: ", environment)
	} else {
		println("No environment specified, defaulting to local")
		environment = "local"
	}

	// Load environment variables from the specified file
	envFile := "config/user." + environment + ".env"
	env.LoadEnv(envFile)
	log.Println("Environment variables loaded from file:", envFile)

	// Connect to the database
	log.Println("Connecting to the database...")
	dsn := "host=" + env.GetEnv("USER_DB_HOST", "localhost") + " user=" + env.GetEnv("USER_DB_USER", "root") + " password=" + env.GetEnv("USER_DB_PASSWORD", "") + " dbname=" + env.GetEnv("USER_DB_NAME", "user_service") + " port=" + env.GetEnv("USER_DB_PORT", "5432") + " sslmode=disable"
	db, err := database.ConnectDB(dsn)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Migrate the database
	log.Println("Migrating the database...")
	err = database.MigrateDB(db)
	if err != nil {
		panic("Failed to migrate the database: " + err.Error())
	}
	
	// Connect to Redis
	log.Println("Connecting to Redis...")
	redisAddr := env.GetEnv("REDIS_HOST", "localhost") + ":" + env.GetEnv("REDIS_PORT", "6379")
	redisClient := redis.NewRedisClient(redisAddr)

	// Initialize the repository, service, and handler
	userRepo := repository.NewUserRepository(db, redisClient)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Start the gRPC server
	log.Println("Starting gRPC server at port", env.GetEnv("USER_SERVER_PORT", "50051"))
	grpc.StartGRPCServer(userHandler, env.GetEnv("USER_SERVER_PORT", "50051"))
}