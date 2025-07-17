package handlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"main/internal/server/interfaces"
	"main/internal/server/models"
	pb "main/proto"
)

type CardsHandler struct {
	pb.UnimplementedCardsServer
	s interfaces.CardsService
	j interfaces.JWTService
}

func NewCardsHandler(s interfaces.CardsService, j interfaces.JWTService) *CardsHandler {
	return &CardsHandler{
		s: s,
		j: j,
	}
}

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

func (h *CardsHandler) Add(ctx context.Context, in *pb.CardCreateRequest) (*pb.CardResponse, error) {
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

	return &pb.CardResponse{
		Id:         result.ID,
		Title:      result.Title,
		Bank:       string(result.Bank),
		Number:     string(result.Number),
		DataEnd:    string(result.DataEnd),
		SecretCode: string(result.SecretCode),
	}, nil
}

func (h *CardsHandler) Update(ctx context.Context, in *pb.CardUpdateRequest) (*pb.CardResponse, error) {
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

	return &pb.CardResponse{
		Id:         result.ID,
		Title:      result.Title,
		Bank:       string(result.Bank),
		Number:     string(result.Number),
		DataEnd:    string(result.DataEnd),
		SecretCode: string(result.SecretCode),
	}, nil
}

func (h *CardsHandler) Delete(ctx context.Context, in *pb.CardRequest) (*emptypb.Empty, error) {
	userID := ctx.Value("userID").(int64)

	err := h.s.Delete(ctx, in.Title, userID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
