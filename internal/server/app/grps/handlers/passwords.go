package handlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"main/internal/server/interfaces"
	pb "main/proto"
)

type PasswordsHandler struct {
	pb.UnimplementedPasswordsServer
	s interfaces.PasswordsService
}

func NewPasswordsHandler(s interfaces.PasswordsService) *PasswordsHandler {
	return &PasswordsHandler{
		s: s,
	}
}

func (h *PasswordsHandler) Get(ctx context.Context, in *pb.PasswordRequest) (*pb.PasswordResponse, error) {
	result, err := h.s.Get(ctx, title, UserID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (h *PasswordsHandler) Add(ctx context.Context, in *pb.PasswordCreateRequest) (*pb.PasswordResponse, error) {
	result, err := h.s.Add(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (h *PasswordsHandler) Update(ctx context.Context, in *pb.PasswordUpdateRequest) (*pb.PasswordResponse, error) {
	result, err := h.s.Update(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (h *PasswordsHandler) Delete(ctx context.Context, in *pb.PasswordRequest) (*emptypb.Empty, error) {
	err := h.s.Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
