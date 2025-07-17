package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"main/internal/server/interfaces"
)

var (
	ErrRPCInvalidToken = status.Errorf(codes.Unauthenticated, "invalid token")
)

func AuthInterceptor(j interfaces.JWTService) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		skipMethods := map[string]bool{
			"/proto.Users/Login":    true,
			"/proto.Users/Register": true,
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
