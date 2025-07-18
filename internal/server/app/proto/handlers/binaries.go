package handlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"main/internal/server/interfaces"
	"main/internal/server/models"
	pb "main/proto"
)

type BinariesHandler struct {
	pb.UnimplementedBinariesServer
	s interfaces.BinariesService
	j interfaces.JWTService
}

func NewBinariesHandler(s interfaces.BinariesService, j interfaces.JWTService) *BinariesHandler {
	return &BinariesHandler{
		s: s,
		j: j,
	}
}

func (h *BinariesHandler) Get(ctx context.Context, in *pb.BinariesRequest) (*pb.BinariesResponse, error) {
	userID := ctx.Value("userID").(int64)

	result, err := h.s.Get(ctx, in.Title, userID)
	if err != nil {
		return nil, err
	}
	return &pb.BinariesResponse{
		Id:    result.ID,
		Title: result.Title,
		Data:  result.Data,
	}, nil
}

func (h *BinariesHandler) Add(ctx context.Context, in *pb.BinariesCreateRequest) (*pb.BinariesShortResponse, error) {
	userID := ctx.Value("userID").(int64)

	cond := models.BinaryData{
		UserID: userID,
		Title:  in.Title,
		Data:   in.Data,
	}

	result, err := h.s.Add(ctx, cond)
	if err != nil {
		return nil, err
	}

	return &pb.BinariesShortResponse{
		Title: result,
	}, nil
}

func (h *BinariesHandler) Update(ctx context.Context, in *pb.BinariesUpdateRequest) (*pb.BinariesShortResponse, error) {
	userID := ctx.Value("userID").(int64)

	cond := models.BinaryData{
		UserID: userID,
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

func (h *BinariesHandler) Delete(ctx context.Context, in *pb.BinariesRequest) (*emptypb.Empty, error) {
	userID := ctx.Value("userID").(int64)

	err := h.s.Delete(ctx, in.Title, userID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
