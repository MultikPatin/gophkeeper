package proto

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"main/internal/server/app/proto/handlers"
	"main/internal/server/app/proto/interceptors"
	pb "main/proto"
)

func NewServer(s *Services, l *zap.SugaredLogger) (*grpc.Server, error) {
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors.NewLoggerInterceptor(l)),
	)

	pb.RegisterUsersServer(srv, handlers.NewUsersHandler(s.users))
	pb.RegisterBinariesServer(srv, handlers.NewBinariesHandler(s.binaries))
	pb.RegisterPasswordsServer(srv, handlers.NewPasswordsHandler(s.passwords))
	pb.RegisterCardsServer(srv, handlers.NewCardsHandler(s.cards))

	return srv, nil
}
