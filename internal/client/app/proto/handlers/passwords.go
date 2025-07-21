package handlers

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "main/proto"
)

type PasswordsHandler struct {
	client pb.PasswordsClient
}

func NewPasswordsHandler(conn *grpc.ClientConn) *PasswordsHandler {
	return &PasswordsHandler{
		client: pb.NewPasswordsClient(conn),
	}
}

func (h *PasswordsHandler) Get(ctx context.Context, in *pb.PasswordRequest) (*pb.PasswordResponse, error) {
	fmt.Println("Getting password")
	return &pb.PasswordResponse{}, nil
}

func (h *PasswordsHandler) Add(ctx context.Context, in *pb.PasswordCreateRequest) (*pb.PasswordShortResponse, error) {
	fmt.Println("Adding password")
	return &pb.PasswordShortResponse{}, nil
}

func (h *PasswordsHandler) Update(ctx context.Context, in *pb.PasswordUpdateRequest) (*pb.PasswordShortResponse, error) {
	fmt.Println("Updating password")
	return &pb.PasswordShortResponse{}, nil
}

func (h *PasswordsHandler) Delete(ctx context.Context, in *pb.PasswordRequest) (*emptypb.Empty, error) {
	fmt.Println("Deleting password")
	return &emptypb.Empty{}, nil
}
