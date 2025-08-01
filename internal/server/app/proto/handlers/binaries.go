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

// BinariesHandler implements the gRPC service definition for managing binary data.
// It delegates requests to the underlying BinariesService for actual business logic execution.
type BinariesHandler struct {
	pb.UnimplementedBinariesServer                            // Base implementation for protobuf-defined gRPC server.
	s                              interfaces.BinariesService // Service for handling binary data operations.
	j                              interfaces.JWTService      // JWT service for authentication purposes.
}

// NewBinariesHandler creates a new instance of BinariesHandler with injected dependencies.
func NewBinariesHandler(s interfaces.BinariesService, j interfaces.JWTService) *BinariesHandler {
	return &BinariesHandler{
		s: s,
		j: j,
	}
}

// Get retrieves a binary data item by title and user ID.
// It extracts the user ID from the context and passes control to the BinariesService.
// Possible errors:
// - ErrBinaryNotFound: If no password matches the given title and user ID.
// - Internal server error if any other issue occurs during processing.
func (h *BinariesHandler) Get(ctx context.Context, in *pb.BinariesRequest) (*pb.BinariesResponse, error) {
	userID := ctx.Value("userID").(int64)

	result, err := h.s.Get(ctx, in.Title, userID)
	if err != nil {
		if errors.Is(err, services.ErrBinaryNotFound) {
			return nil, status.Errorf(codes.NotFound, "binary with title '%s' was not found.", in.Title)
		}
		return nil, status.Error(codes.Internal, "Internal server error.")
	}

	return &pb.BinariesResponse{
		Id:    result.ID,
		Title: result.Title,
		Data:  result.Data,
	}, nil
}

// Add creates a new binary data entry.
// It populates a BinaryData model and invokes the BinariesService to perform the insertion.
// Possible errors:
// - ErrBinaryAlreadyExists: If a password with the same title already exists for this user.
// - Internal server error if any other issue occurs during processing.
func (h *BinariesHandler) Add(ctx context.Context, in *pb.BinariesCreateRequest) (*pb.BinariesShortResponse, error) {
	userID := ctx.Value("userID").(int64)

	cond := models.BinaryData{
		UserID: userID,
		Title:  in.Title,
		Data:   in.Data,
	}

	result, err := h.s.Add(ctx, cond)
	if err != nil {
		if errors.Is(err, services.ErrBinaryAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "A binary with title '%s' already exists.", in.Title)
		}
		return nil, status.Error(codes.Internal, "Internal server error.")
	}

	return &pb.BinariesShortResponse{
		Title: result,
	}, nil
}

// Update modifies an existing binary data entry.
// It prepares a BinaryData model and triggers the BinariesService to execute the update.
// Possible errors:
// - Internal server error if any issue occurs during processing.
func (h *BinariesHandler) Update(ctx context.Context, in *pb.BinariesUpdateRequest) (*pb.BinariesShortResponse, error) {
	userID := ctx.Value("userID").(int64)

	cond := models.BinaryData{
		UserID: userID,
		Title:  in.Title,
		Data:   in.Data,
	}

	result, err := h.s.Update(ctx, cond)
	if err != nil {
		return nil, err
	}

	return &pb.BinariesShortResponse{
		Title: result,
	}, nil
}

// Delete removes a binary data entry by title and user ID.
// It extracts the user ID from the context and forwards the removal request to the BinariesService.
// Possible errors:
// - Internal server error if any issue occurs during processing.
func (h *BinariesHandler) Delete(ctx context.Context, in *pb.BinariesRequest) (*emptypb.Empty, error) {
	userID := ctx.Value("userID").(int64)

	err := h.s.Delete(ctx, in.Title, userID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
