package handlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"main/internal/server/interfaces"
	pb "main/proto"
)

type CardsHandler struct {
	pb.UnimplementedCardsServer
	s interfaces.CardsService
}

func NewCardsHandler(s interfaces.CardsService) *CardsHandler {
	return &CardsHandler{
		s: s,
	}
}

func (h *CardsHandler) Get(ctx context.Context, in *pb.CardRequest) (*pb.CardResponse, error) {
	result, err := h.s.Get(ctx, title, UserID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (h *CardsHandler) Add(ctx context.Context, in *pb.CardCreateRequest) (*pb.CardResponse, error) {
	result, err := h.s.Add(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (h *CardsHandler) Update(ctx context.Context, in *pb.CardUpdateRequest) (*pb.CardResponse, error) {
	result, err := h.s.Update(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (h *CardsHandler) Delete(ctx context.Context, in *pb.CardRequest) (*emptypb.Empty, error) {
	err := h.s.Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
