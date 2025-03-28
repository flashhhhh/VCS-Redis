package grpc

import (
	"log"
	"net"
	"user_service/internal/handler"
	"user_service/pb"

	"google.golang.org/grpc"
)

func StartGRPCServer(userHandler *handler.UserHandler) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userHandler)

	log.Println("gRPC server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}