package proto

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"main/internal/server/app/proto/handlers"
	"main/internal/server/app/proto/interceptors"
	"main/internal/server/auth"
	"main/internal/server/config"
	pb "main/proto"
)

func NewServer(s *Services, c *config.Config, l *zap.SugaredLogger) (*grpc.Server, error) {
	jwtService, err := auth.NewJWTService(c.JWTSecret, c.JWTExpiration)
	if err != nil {
		return nil, err
	}

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors.NewLoggerInterceptor(l)),
	)

	pb.RegisterUsersServer(srv, handlers.NewUsersHandler(s.users, jwtService))
	pb.RegisterBinariesServer(srv, handlers.NewBinariesHandler(s.binaries, jwtService))
	pb.RegisterPasswordsServer(srv, handlers.NewPasswordsHandler(s.passwords, jwtService))
	pb.RegisterCardsServer(srv, handlers.NewCardsHandler(s.cards, jwtService))

	return srv, nil
}
