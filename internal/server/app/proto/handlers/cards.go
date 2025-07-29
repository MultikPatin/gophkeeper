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

// CardsHandler implements the gRPC service definition for managing credit card data.
// It delegates requests to the underlying CardsService for actual business logic execution.
type CardsHandler struct {
	pb.UnimplementedCardsServer                         // Base implementation for protobuf-defined gRPC server.
	s                           interfaces.CardsService // Service for handling credit card operations.
	j                           interfaces.JWTService   // JWT service for authentication purposes.
}

// NewCardsHandler creates a new instance of CardsHandler with injected dependencies.
func NewCardsHandler(s interfaces.CardsService, j interfaces.JWTService) *CardsHandler {
	return &CardsHandler{
		s: s,
		j: j,
	}
}

// Get retrieves a credit card by title and user ID.
// It extracts the user ID from the context and passes control to the CardsService.
// Possible errors:
// - ErrCardNotFound: If no password matches the given title and user ID.
// - Internal server error if any other issue occurs during processing.
func (h *CardsHandler) Get(ctx context.Context, in *pb.CardRequest) (*pb.CardResponse, error) {
	userID := ctx.Value("userID").(int64)

	result, err := h.s.Get(ctx, in.Title, userID)
	if err != nil {
		if errors.Is(err, services.ErrCardNotFound) {
			return nil, status.Errorf(codes.NotFound, "card with title '%s' was not found.", in.Title)
		}
		return nil, status.Error(codes.Internal, "Internal server error.")
	}

	return &pb.CardResponse{
		Id:         result.ID,
		Title:      result.Title,
		Bank:       string(result.Bank),
		Number:     string(result.Number),
		DataEnd:    string(result.DataEnd),
		SecretCode: string(result.SecretCode),
	}, nil
}

// Add creates a new credit card entry.
// It populates a Card model and invokes the CardsService to perform the insertion.
// Possible errors:
// - ErrCardAlreadyExists: If a password with the same title already exists for this user.
// - Internal server error if any other issue occurs during processing.
func (h *CardsHandler) Add(ctx context.Context, in *pb.CardCreateRequest) (*pb.CardShortResponse, error) {
	userID := ctx.Value("userID").(int64)

	cond := models.Card{
		UserID:     userID,
		Title:      in.Title,
		Bank:       []byte(in.Bank),
		Number:     []byte(in.Number),
		DataEnd:    []byte(in.DataEnd),
		SecretCode: []byte(in.SecretCode),
	}

	result, err := h.s.Add(ctx, cond)
	if err != nil {
		if errors.Is(err, services.ErrCardAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "A card with title '%s' already exists.", in.Title)
		}
		return nil, status.Error(codes.Internal, "Internal server error.")
	}

	return &pb.CardShortResponse{
		Title: result,
	}, nil
}

// Update modifies an existing credit card entry.
// It prepares a Card model and triggers the CardsService to execute the update.
// Possible errors:
// - Internal server error if any issue occurs during processing.
func (h *CardsHandler) Update(ctx context.Context, in *pb.CardUpdateRequest) (*pb.CardShortResponse, error) {
	userID := ctx.Value("userID").(int64)

	cond := models.Card{
		UserID:     userID,
		Title:      in.Title,
		Bank:       []byte(in.Bank),
		Number:     []byte(in.Number),
		DataEnd:    []byte(in.DataEnd),
		SecretCode: []byte(in.SecretCode),
	}

	result, err := h.s.Update(ctx, cond)
	if err != nil {
		return nil, err
	}

	return &pb.CardShortResponse{
		Title: result,
	}, nil
}

// Delete removes a credit card entry by title and user ID.
// It extracts the user ID from the context and forwards the removal request to the CardsService.
// Possible errors:
// - Internal server error if any issue occurs during processing.
func (h *CardsHandler) Delete(ctx context.Context, in *pb.CardRequest) (*emptypb.Empty, error) {
	userID := ctx.Value("userID").(int64)

	err := h.s.Delete(ctx, in.Title, userID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
