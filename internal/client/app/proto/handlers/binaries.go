package handlers

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "main/proto"
)

type BinariesHandler struct {
	client pb.BinariesClient
}

func NewBinariesHandler(conn *grpc.ClientConn) *BinariesHandler {
	return &BinariesHandler{
		client: pb.NewBinariesClient(conn),
	}
}

func (h *BinariesHandler) Get(ctx context.Context, in *pb.BinariesRequest) (*pb.BinariesResponse, error) {
	fmt.Println("Getting binaries...")
	return &pb.BinariesResponse{}, nil
}

func (h *BinariesHandler) Add(ctx context.Context, in *pb.BinariesCreateRequest) (*pb.BinariesShortResponse, error) {
	fmt.Println("Adding binaries...")
	return &pb.BinariesShortResponse{}, nil
}

func (h *BinariesHandler) Update(ctx context.Context, in *pb.BinariesUpdateRequest) (*pb.BinariesShortResponse, error) {
	fmt.Println("Updating binaries...")
	return &pb.BinariesShortResponse{}, nil
}

func (h *BinariesHandler) Delete(ctx context.Context, in *pb.BinariesRequest) (*emptypb.Empty, error) {
	fmt.Println("Deleting binaries...")
	return nil, nil
}
