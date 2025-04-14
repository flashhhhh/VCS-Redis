package client

import (
	"context"
	"log"
	"gateway/pb"

	"google.golang.org/grpc"
)

type UserClient interface {
    CreateUser(username, password, name string) (*pb.EmptyResponse, error)
    Login(username, password string) (*pb.LoginResponse, error)
    GetUserByID(id int32) (*pb.UserResponse, error)
    GetAllUsers() (*pb.UsersResponse, error)
}

type userClient struct {
    UserService pb.UserServiceClient
}

func NewUserClient(userServiceAddress string) UserClient {
    conn, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect to user service: %v", err)
    }
    return &userClient{
        UserService: pb.NewUserServiceClient(conn),
    }
}

func (user *userClient) CreateUser(username, password, name string) (*pb.EmptyResponse, error) {
    req := &pb.CreateUserRequest{
        Username: username,
        Password: password,
        Name:     name,
    }
    resp, err := user.UserService.CreateUser(context.Background(), req)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

func (user *userClient) Login(username, password string) (*pb.LoginResponse, error) {
    req := &pb.LoginRequest{
        Username: username,
        Password: password,
    }
    resp, err := user.UserService.Login(context.Background(), req)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

func (user *userClient) GetUserByID(id int32) (*pb.UserResponse, error) {
    req := &pb.GetUserByIDRequest{
        Id: id,
    }
    resp, err := user.UserService.GetUserByID(context.Background(), req)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

func (user *userClient) GetAllUsers() (*pb.UsersResponse, error) {
    req := &pb.EmptyRequest{}
    resp, err := user.UserService.GetAllUsers(context.Background(), req)
    if err != nil {
        return nil, err
    }
    return resp, nil
}