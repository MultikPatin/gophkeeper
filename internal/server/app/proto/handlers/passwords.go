package handlers

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"main/internal/server/interfaces"
	"main/internal/server/models"
	"main/internal/server/services"
	pb "main/proto"
)

// PasswordsHandler implements the gRPC service definition for managing password data.
// It delegates requests to the underlying PasswordsService for actual business logic execution.
type PasswordsHandler struct {
	pb.UnimplementedPasswordsServer                             // Base implementation for protobuf-defined gRPC server.
	s                               interfaces.PasswordsService // Service for handling password operations.
	j                               interfaces.JWTService       // JWT service for authentication purposes.
}

// NewPasswordsHandler creates a new instance of PasswordsHandler with injected dependencies.
func NewPasswordsHandler(s interfaces.PasswordsService, j interfaces.JWTService) *PasswordsHandler {
	return &PasswordsHandler{
		s: s,
		j: j,
	}
}

// Get retrieves a password by title and user ID.
// It extracts the user ID from the context and passes control to the PasswordsService.
// Possible errors:
// - ErrPasswordNotFound: If no password matches the given title and user ID.
// - Internal server error if any other issue occurs during processing.
func (h *PasswordsHandler) Get(ctx context.Context, in *pb.PasswordRequest) (*pb.PasswordResponse, error) {
	userID := ctx.Value("userID").(int64)

	result, err := h.s.Get(ctx, in.Title, userID)
	if err != nil {
		if errors.Is(err, services.ErrPasswordNotFound) {
			return nil, status.Errorf(codes.NotFound, "password with title '%s' was not found.", in.Title)
		}
		return nil, status.Error(codes.Internal, "Internal server error.")
	}

	return &pb.PasswordResponse{
		Id:       result.ID,
		Title:    result.Title,
		Login:    string(result.Login),
		Password: string(result.Password),
	}, nil
}

// Add creates a new password entry.
// It populates a Password model and invokes the PasswordsService to perform the insertion.
// Possible errors:
// - ErrPasswordAlreadyExists: If a password with the same title already exists for this user.
// - Internal server error if any other issue occurs during processing.
func (h *PasswordsHandler) Add(ctx context.Context, in *pb.PasswordCreateRequest) (*pb.PasswordShortResponse, error) {
	userID := ctx.Value("userID").(int64)

	cond := models.Password{
		UserID:   userID,
		Title:    in.Title,
		Login:    []byte(in.Login),
		Password: []byte(in.Password),
	}

	result, err := h.s.Add(ctx, cond)
	if err != nil {
		if errors.Is(err, services.ErrPasswordAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "A password with title '%s' already exists.", in.Title)
		}
		return nil, status.Error(codes.Internal, "Internal server error.")
	}

	return &pb.PasswordShortResponse{
		Title: result,
	}, nil
}

// Update modifies an existing password entry.
// It prepares a Password model and triggers the PasswordsService to execute the update.
// Possible errors:
// - Internal server error if any issue occurs during processing.
func (h *PasswordsHandler) Update(ctx context.Context, in *pb.PasswordUpdateRequest) (*pb.PasswordShortResponse, error) {
	userID := ctx.Value("userID").(int64)

	cond := models.Password{
		UserID:   userID,
		Title:    in.Title,
		Login:    []byte(in.Login),
		Password: []byte(in.Password),
	}

	result, err := h.s.Update(ctx, cond)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error.")
	}

	return &pb.PasswordShortResponse{
		Title: result,
	}, nil
}

// Delete removes a password entry by title and user ID.
// It extracts the user ID from the context and forwards the removal request to the PasswordsService.
// Possible errors:
// - Internal server error if any issue occurs during processing.
func (h *PasswordsHandler) Delete(ctx context.Context, in *pb.PasswordRequest) (*emptypb.Empty, error) {
	userID := ctx.Value("userID").(int64)

	err := h.s.Delete(ctx, in.Title, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error.")
	}
	return &emptypb.Empty{}, nil
}
