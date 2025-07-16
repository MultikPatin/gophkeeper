package grps

import (
	"google.golang.org/grpc"
	"main/internal/server/app/grps/handlers"
	pb "main/proto"
)

func NewServer(s *Services) (*grpc.Server, error) {
	srv := grpc.NewServer()

	pb.RegisterUsersServer(srv, handlers.NewUsersHandler(s.users))
	pb.RegisterBinariesServer(srv, handlers.NewBinariesHandler(s.binaries))
	pb.RegisterPasswordsServer(srv, handlers.NewPasswordsHandler(s.passwords))
	pb.RegisterCardsServer(srv, handlers.NewCardsHandler(s.cards))

	return srv, nil
}
