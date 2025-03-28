package handler

import (
	"context"
	"user_service/internal/service"
	"user_service/pb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.EmptyResponse, error) {
	err := h.service.CreateUser(ctx, req.Username, req.Password, req.Name)
	if err != nil {
		return &pb.EmptyResponse{}, err
	}
	return &pb.EmptyResponse{}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := h.service.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	
	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (h *UserHandler) GetUserByID(ctx context.Context, req *pb.IDRequest) (*pb.UserResponse, error) {
	user, err := h.service.GetUserByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:       int32(user.ID),
		Username: user.Username,
		Name:     user.Name,
	}, nil
}

func (h *UserHandler) GetUserByUsername(ctx context.Context, req *pb.UsernameRequest) (*pb.UserResponse, error) {
	user, err := h.service.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:       int32(user.ID),
		Username: user.Username,
		Name:     user.Name,
	}, nil
}

func (h *UserHandler) GetAllUsers(ctx context.Context, req *pb.EmptyRequest) (*pb.UserArrayResponse, error) {
	users, err := h.service.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	var userResponses []*pb.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, &pb.UserResponse{
			Id:       int32(user.ID),
			Username: user.Username,
			Name:     user.Name,
		})
	}
	return &pb.UserArrayResponse{Users: userResponses}, nil
}