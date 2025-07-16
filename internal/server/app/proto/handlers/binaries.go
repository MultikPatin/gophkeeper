package handlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"main/internal/server/interfaces"
	pb "main/proto"
)

type BinariesHandler struct {
	pb.UnimplementedBinariesServer
	s interfaces.BinariesService
}

func NewBinariesHandler(s interfaces.BinariesService) *BinariesHandler {
	return &BinariesHandler{
		s: s,
	}
}

func (h *BinariesHandler) Get(ctx context.Context, in *pb.BinariesRequest) (*pb.BinariesResponse, error) {
	result, err := h.s.Get(ctx, title, UserID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (h *BinariesHandler) Add(ctx context.Context, in *pb.BinariesCreateRequest) (*pb.BinariesResponse, error) {
	result, err := h.s.Add(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (h *BinariesHandler) Update(ctx context.Context, in *pb.BinariesUpdateRequest) (*pb.BinariesResponse, error) {
	result, err := h.s.Update(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (h *BinariesHandler) Delete(ctx context.Context, in *pb.BinariesRequest) (*emptypb.Empty, error) {
	err := h.s.Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
