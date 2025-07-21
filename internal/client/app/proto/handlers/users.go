package handlers

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "main/proto"
)

type UsersHandler struct {
	client pb.UsersClient
}

func NewUsersHandler(conn *grpc.ClientConn) *UsersHandler {
	return &UsersHandler{
		client: pb.NewUsersClient(conn),
	}
}

func (h *UsersHandler) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	fmt.Println("Registering user...")
	return &pb.RegisterResponse{
		Token: "1234567890",
	}, nil
}

func (h *UsersHandler) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println("Logging in user...")
	return &pb.LoginResponse{}, nil
}
