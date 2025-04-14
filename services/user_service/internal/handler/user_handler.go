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
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.EmptyResponse, error) {
	// Call the service to create a user
	err := h.service.CreateUser(ctx, req.Username, req.Password, req.Name)
	if err != nil {
		return nil, err
	}

	// Return an empty response
	return &pb.EmptyResponse{}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Call the service to login
	token, err := h.service.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// Return the user information
	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (h *UserHandler) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.UserResponse, error) {
	// Call the service to get user by ID
	user, err := h.service.GetUserByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	// Return the user information
	return &pb.UserResponse{
		Id:       int32(user.ID),
		Username: user.Username,
		Name:     user.Name,
	}, nil
}

func (h *UserHandler) GetAllUsers(ctx context.Context, req *pb.EmptyRequest) (*pb.UsersResponse, error) {
	// Call the service to get all users
	users, err := h.service.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the users to the response format
	var userResponses []*pb.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, &pb.UserResponse{
			Id:       int32(user.ID),
			Username: user.Username,
			Name:     user.Name,
		})
	}

	// Return the list of users
	return &pb.UsersResponse{
		Users: userResponses,
	}, nil
}