package handlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"main/internal/server/interfaces"
	"main/internal/server/models"
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
func (h *PasswordsHandler) Get(ctx context.Context, in *pb.PasswordRequest) (*pb.PasswordResponse, error) {
	userID := ctx.Value("userID").(int64)

	result, err := h.s.Get(ctx, in.Title, userID)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &pb.PasswordShortResponse{
		Title: result,
	}, nil
}

// Update modifies an existing password entry.
// It prepares a Password model and triggers the PasswordsService to execute the update.
func (h *PasswordsHandler) Update(ctx context.Context, in *pb.PasswordUpdateRequest) (*pb.PasswordShortResponse, error) {
	userID := ctx.Value("userID").(int64)

	cond := models.Password{
		UserID:   userID,
		Login:    []byte(in.Login),
		Password: []byte(in.Password),
	}

	result, err := h.s.Update(ctx, cond)
	if err != nil {
		return nil, err
	}

	return &pb.PasswordShortResponse{
		Title: result,
	}, nil
}

// Delete removes a password entry by title and user ID.
// It extracts the user ID from the context and forwards the removal request to the PasswordsService.
func (h *PasswordsHandler) Delete(ctx context.Context, in *pb.PasswordRequest) (*emptypb.Empty, error) {
	userID := ctx.Value("userID").(int64)

	err := h.s.Delete(ctx, in.Title, userID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
