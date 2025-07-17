package helpers

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"main/internal/server/interfaces"
)

var (
	ErrRPCInvalidToken = status.Errorf(codes.Unauthenticated, "invalid token")
)

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
