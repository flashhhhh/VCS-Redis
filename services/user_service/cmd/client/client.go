package main

import (
	"context"
	"fmt"
	"user_service/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	// Example usage of the client
	// CreateUser
	_, err = client.CreateUser(context.Background(), &pb.CreateUserRequest{
		Username: "testuser",
		Password: "password",
		Name:     "Test User",
	})
	if err != nil {
		panic(err)
	}

 	// Login
	resp, _ := client.Login(context.Background(), &pb.LoginRequest{
		Username: "testuser",
		Password: "password",
	})
	fmt.Println(resp)

	// Get all users
	users, _ := client.GetAllUsers(context.Background(), &pb.EmptyRequest{})
	for _, user := range users.Users {
		fmt.Printf("User: %s, ID: %d\n", user.Username, user.Id)
	}
}