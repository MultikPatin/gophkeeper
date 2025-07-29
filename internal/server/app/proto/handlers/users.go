package handlers

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"main/internal/server/interfaces"
	"main/internal/server/models"
	"main/internal/server/services"
	pb "main/proto"
	"time"
)

// UsersHandler implements the gRPC service definition for user management.
// Delegates requests to the underlying UsersService for actual business logic execution.
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

// Register handles user registration by delegating to the UsersService and generates a JWT token on successful completion.
// Possible errors:
// - ErrLoginAlreadyExists: The provided login is already taken by another user.
// - Internal server error if registration fails.
func (h *UsersHandler) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cond := models.User{
		Login:    in.Login,
		Password: in.Password,
	}

	userID, err := h.s.Register(ctx, cond)
	if err != nil {
		if errors.Is(err, services.ErrLoginAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "A user with login '%s' already exists.", in.Login)
		}
		return nil, status.Error(codes.Internal, "Error registering new user.")
	}

	token, err := h.j.Generate(userID)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Token: token,
	}, nil
}

// Login authenticates a user and returns a JWT token upon successful authentication.
// Possible errors:
// - ErrUserNotFound: User with specified login does not exist.
// - ErrInvalidCredentials: Provided username or password is incorrect.
// - Internal server error when authentication fails.
func (h *UsersHandler) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	cond := models.User{
		Login:    in.Login,
		Password: in.Password,
	}

	userID, err := h.s.Login(ctx, cond)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "User with login '%s' was not found.", in.Login)
		}
		if errors.Is(err, services.ErrInvalidCredentials) {
			return nil, status.Errorf(codes.Unauthenticated, "Provided credentials are invalid.")
		}
		return nil, status.Error(codes.Internal, "Error logging into account.")
	}

	token, err := h.j.Generate(userID)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}
