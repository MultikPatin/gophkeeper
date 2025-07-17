package handlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"main/internal/server/interfaces"
	"main/internal/server/models"
	pb "main/proto"
)

type PasswordsHandler struct {
	pb.UnimplementedPasswordsServer
	s interfaces.PasswordsService
	j interfaces.JWTService
}

func NewPasswordsHandler(s interfaces.PasswordsService, j interfaces.JWTService) *PasswordsHandler {
	return &PasswordsHandler{
		s: s,
		j: j,
	}
}

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

func (h *PasswordsHandler) Add(ctx context.Context, in *pb.PasswordCreateRequest) (*pb.PasswordResponse, error) {
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

	return &pb.PasswordResponse{
		Id:       result.ID,
		Title:    result.Title,
		Login:    string(result.Login),
		Password: string(result.Password),
	}, nil
}

func (h *PasswordsHandler) Update(ctx context.Context, in *pb.PasswordUpdateRequest) (*pb.PasswordResponse, error) {
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

	return &pb.PasswordResponse{
		Id:       result.ID,
		Title:    result.Title,
		Login:    string(result.Login),
		Password: string(result.Password),
	}, nil
}

func (h *PasswordsHandler) Delete(ctx context.Context, in *pb.PasswordRequest) (*emptypb.Empty, error) {
	userID := ctx.Value("userID").(int64)

	err := h.s.Delete(ctx, in.Title, userID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
