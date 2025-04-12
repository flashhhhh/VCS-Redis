/*
	Package main is the entry point of the user service.
*/

package main

import (
	"fmt"
	"net"

	"user_service/internal/handler"
	pb "user_service/pb"

	"github.com/flashhhhh/pkg/env"
	"google.golang.org/grpc"
)

func main() {
	env.LoadEnv("config/user.env")

	lis, err := net.Listen("tcp", env.GetEnv("USER_PORT", ":50051"))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &handler.Server{})

	fmt.Println("Starting server at " + env.GetEnv("USER_PORT", ":50051"))
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}