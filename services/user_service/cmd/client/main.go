package main

import (
	"context"
	"log"
	"user_service/pb"

	"github.com/flashhhhh/pkg/env"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	grpcServerAddress := env.GetEnv("USER_SERVER_HOST", "localhost") + ":" + env.GetEnv("USER_SERVER_PORT", "50051")

	conn, err := grpc.Dial(grpcServerAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	// Example usage of the client
	
	/*
		Create a user
	*/
	// _, err = client.CreateUser(context.Background(), &pb.CreateUserRequest{
	// 	Username: "testuser2",
	// 	Password: "password",
	// 	Name:     "Test User 2",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("User created successfully!")

	/*
		Log in a user
	*/
	// token, err := client.Login(context.Background(), &pb.LoginRequest{
	// 	Username: "testuser2",
	// 	Password: "password",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("Login successful! Token:", token.Token)

	/*
		Get user by ID
	*/
	// user, err := client.GetUserByID(context.Background(), &pb.GetUserByIDRequest{
	// 	Id: 1,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("User details:", user)

	/*
		Get all users
	*/
	users, err := client.GetAllUsers(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		panic(err)
	}
	log.Println("All users:")
	for _, user := range users.Users {
		log.Printf("ID: %d, Username: %s, Name: %s\n", user.Id, user.Username, user.Name)
	}
}