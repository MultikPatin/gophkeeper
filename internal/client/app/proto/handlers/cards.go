package handlers

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "main/proto"
)

type CardsHandler struct {
	client pb.CardsClient
}

func NewCardsHandler(conn *grpc.ClientConn) *CardsHandler {
	return &CardsHandler{
		client: pb.NewCardsClient(conn),
	}
}

func (h *CardsHandler) Get(ctx context.Context, in *pb.CardRequest) (*pb.CardResponse, error) {
	fmt.Println("Getting card...")
	return &pb.CardResponse{}, nil
}

func (h *CardsHandler) Add(ctx context.Context, in *pb.CardCreateRequest) (*pb.CardShortResponse, error) {
	fmt.Println("Adding card...")
	return &pb.CardShortResponse{}, nil
}

func (h *CardsHandler) Update(ctx context.Context, in *pb.CardUpdateRequest) (*pb.CardShortResponse, error) {
	fmt.Println("Updating card...")
	return &pb.CardShortResponse{}, nil
}

func (h *CardsHandler) Delete(ctx context.Context, in *pb.CardRequest) (*emptypb.Empty, error) {
	fmt.Println("Deleting card...")
	return nil, nil
}
