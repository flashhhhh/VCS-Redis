package handler

import (
	"context"
	"user_service/internal/service"
	pb "user_service/pb"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.EmptyResponse, error) {
	username := in.Username
	password := in.Password
	name := in.Name

	err := service.CreateUser(username, password, name)
	if err != nil {
		return nil, err
	}

	return &pb.EmptyResponse{}, nil
}

func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	username := in.Username
	password := in.Password

	token, err := service.Login(username, password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (s *Server) GetUserByID(ctx context.Context, in *pb.IDRequest) (*pb.UserResponse, error) {
	id := int(in.Id)

	userInfo, err := service.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id: int32(userInfo["id"].(int)),
		Username: userInfo["username"].(string),
		Name: userInfo["name"].(string),
	}, nil
}

func (s *Server) GetUserByName(ctx context.Context, in *pb.UsernameRequest) (*pb.UserResponse, error) {
	username := in.Username

	userInfo, err := service.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id: int32(userInfo["id"].(int)),
		Username: userInfo["username"].(string),
		Name: userInfo["name"].(string),
	}, nil
}

func (s *Server) GetAllUsers(ctx context.Context, in *pb.EmptyRequest) (*pb.UserArrayResponse, error) {
	userArray, err := service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var users []*pb.UserResponse
	for _, user := range userArray {
		users = append(users, &pb.UserResponse{
			Id: int32(user["id"].(int)),
			Username: user["username"].(string),
			Name: user["name"].(string),
		})
	}

	return &pb.UserArrayResponse{
		Users: users,
	}, nil
}