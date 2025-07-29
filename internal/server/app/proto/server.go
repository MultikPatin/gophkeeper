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

// NewServer initializes and configures a gRPC server instance.
// It incorporates interceptors for logging and authentication, and registers handlers for gRPC services.
func NewServer(s *Services, c *config.Config, l *zap.SugaredLogger) (*grpc.Server, error) {
	// Initialize JWT service for authentication purposes.
	j, err := auth.NewJWTService(c.JWTSecret, c.JWTExpiration)
	if err != nil {
		return nil, err
	}

	// Instantiate a new gRPC server with chained interceptors for logging and authentication.
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.LoggerInterceptor(l), // Logging interceptor.
			interceptors.AuthInterceptor(j),   // Authentication interceptor.
		),
	)

	// Register gRPC service handlers for respective domains.
	pb.RegisterUsersServer(srv, handlers.NewUsersHandler(s.users, j))             // Handler for user-related RPCs.
	pb.RegisterBinariesServer(srv, handlers.NewBinariesHandler(s.binaries, j))    // Handler for binary data-related RPCs.
	pb.RegisterPasswordsServer(srv, handlers.NewPasswordsHandler(s.passwords, j)) // Handler for password-related RPCs.
	pb.RegisterCardsServer(srv, handlers.NewCardsHandler(s.cards, j))             // Handler for credit card-related RPCs.

	return srv, nil
}
