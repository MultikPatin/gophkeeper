package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"main/internal/server/interfaces"
)

// ErrRPCInvalidToken represents an error when the provided JWT token is invalid or missing.
var (
	ErrRPCInvalidToken = status.Errorf(codes.Unauthenticated, "invalid token")
)

// AuthInterceptor is a gRPC Unary Server Interceptor that enforces authentication.
// It extracts the JWT token from the request metadata and verifies it using the JWTService.
// If the token is valid, the user ID is propagated through the context for downstream handlers.
func AuthInterceptor(j interfaces.JWTService) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		skipMethods := map[string]bool{
			"/proto.Users/Login":    true, // Allows login requests without authentication.
			"/proto.Users/Register": true, // Allows registration requests without authentication.
		}

		if skipMethods[info.FullMethod] {
			return handler(ctx, req)
		}

		userID, err := GetUserIDFromMD(ctx, j)
		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, "userID", userID)
		return handler(ctx, req)
	}
}

// GetUserIDFromMD retrieves the JWT token from the request metadata and verifies it.
// If the token is successfully validated, the associated user ID is returned.
func GetUserIDFromMD(ctx context.Context, j interfaces.JWTService) (int64, error) {
	var token string

	if md, ok := metadata.FromIncomingContext(ctx); !ok {
		return -1, ErrRPCInvalidToken
	} else if vals := md.Get("token"); len(vals) > 0 && vals[0] != "" {
		token = vals[0]
	} else {
		return -1, ErrRPCInvalidToken
	}

	userID, err := j.Verify(token)
	if err != nil {
		return -1, ErrRPCInvalidToken
	}
	return userID, nil
}
