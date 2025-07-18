package handlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"main/internal/server/interfaces"
	"main/internal/server/models"
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
func (h *CardsHandler) Get(ctx context.Context, in *pb.CardRequest) (*pb.CardResponse, error) {
	userID := ctx.Value("userID").(int64)

	result, err := h.s.Get(ctx, in.Title, userID)
	if err != nil {
		return nil, err
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
func (h *CardsHandler) Add(ctx context.Context, in *pb.CardCreateRequest) (*pb.CardShortResponse, error) {
	userID := ctx.Value("userID").(int64)

	cond := models.Card{
		Title:      in.Title,
		UserID:     userID,
		Bank:       []byte(in.Bank),
		Number:     []byte(in.Number),
		DataEnd:    []byte(in.DataEnd),
		SecretCode: []byte(in.SecretCode),
	}

	result, err := h.s.Add(ctx, cond)
	if err != nil {
		return nil, err
	}

	return &pb.CardShortResponse{
		Title: result,
	}, nil
}

// Update modifies an existing credit card entry.
// It prepares a Card model and triggers the CardsService to execute the update.
func (h *CardsHandler) Update(ctx context.Context, in *pb.CardUpdateRequest) (*pb.CardShortResponse, error) {
	userID := ctx.Value("userID").(int64)

	cond := models.Card{
		UserID:     userID,
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
func (h *CardsHandler) Delete(ctx context.Context, in *pb.CardRequest) (*emptypb.Empty, error) {
	userID := ctx.Value("userID").(int64)

	err := h.s.Delete(ctx, in.Title, userID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
