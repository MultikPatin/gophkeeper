package handlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"main/internal/server/app/proto/helpers"
	"main/internal/server/interfaces"
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
	UserID, err := helpers.GetUserIDFromMD(ctx, h.j)
	if err != nil {
		return nil, err
	}

	result, err := h.s.Get(ctx, title, UserID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (h *PasswordsHandler) Add(ctx context.Context, in *pb.PasswordCreateRequest) (*pb.PasswordResponse, error) {
	UserID, err := helpers.GetUserIDFromMD(ctx, h.j)
	if err != nil {
		return nil, err
	}

	result, err := h.s.Add(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (h *PasswordsHandler) Update(ctx context.Context, in *pb.PasswordUpdateRequest) (*pb.PasswordResponse, error) {
	UserID, err := helpers.GetUserIDFromMD(ctx, h.j)
	if err != nil {
		return nil, err
	}

	result, err := h.s.Update(ctx, cond)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (h *PasswordsHandler) Delete(ctx context.Context, in *pb.PasswordRequest) (*emptypb.Empty, error) {
	UserID, err := helpers.GetUserIDFromMD(ctx, h.j)
	if err != nil {
		return nil, err
	}

	err := h.s.Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
