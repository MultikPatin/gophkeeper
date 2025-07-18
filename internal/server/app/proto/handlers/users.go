package handlers

import (
	"context"
	"main/internal/server/interfaces"
	"main/internal/server/models"
	pb "main/proto"
)

// UsersHandler implements the gRPC service definition for user management.
// It delegates requests to the underlying UsersService for actual business logic execution.
type UsersHandler struct {
	pb.UnimplementedUsersServer                         // Base implementation for protobuf-defined gRPC server.
	s                           interfaces.UsersService // Service for handling user-related operations.
	j                           interfaces.JWTService   // JWT service for authentication and token generation.
}

// NewUsersHandler creates a new instance of UsersHandler with injected dependencies.
func NewUsersHandler(s interfaces.UsersService, j interfaces.JWTService) *UsersHandler {
	return &UsersHandler{
		s: s,
		j: j,
	}
}

// Register handles user registration by delegating to the UsersService and generating a JWT token upon success.
func (h *UsersHandler) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	cond := models.User{
		Login:    in.Login,
		Password: in.Password,
	}

	userID, err := h.s.Register(ctx, cond)
	if err != nil {
		return nil, err
	}

	token, err := h.j.Generate(userID)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Token: token,
	}, nil
}

// Login authenticates a user and issues a JWT token upon successful login.
func (h *UsersHandler) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	cond := models.User{
		Login:    in.Login,
		Password: in.Password,
	}

	userID, err := h.s.Login(ctx, cond)
	if err != nil {
		return nil, err
	}

	token, err := h.j.Generate(userID)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}
