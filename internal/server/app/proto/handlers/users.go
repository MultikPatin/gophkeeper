package handlers

import (
	"context"
	"main/internal/server/interfaces"
	pb "main/proto"
)

type UsersHandler struct {
	pb.UnimplementedUsersServer
	s interfaces.UsersService
	j interfaces.JWTService
}

func NewUsersHandler(s interfaces.UsersService, j interfaces.JWTService) *UsersHandler {
	return &UsersHandler{
		s: s,
		j: j,
	}
}

func (h *UsersHandler) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	err := h.s.Register(ctx, cond)
	if err != nil {
		return err
	}
	return nil
}

func (h *UsersHandler) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	result, err := h.s.Login(ctx, Login)
	if err != nil {
		return nil, err
	}
	return result, nil
}
