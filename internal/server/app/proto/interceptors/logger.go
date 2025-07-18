package interceptors

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

// LoggerInterceptor crete a gRPC interceptor that logs requests and responses.
func LoggerInterceptor(logger *zap.SugaredLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		start := time.Now()
		md, _ := metadata.FromIncomingContext(ctx)

		logger.Info("request  | ",
			"method: ", info.FullMethod,
			" metadata: ", md)

		logger.Debug("incoming | ",
			"request: ", req)

		resp, err = handler(ctx, req)

		logger.Info("response | ",
			"duration: ", time.Since(start),
			" ok: ", err == nil)

		logger.Debug("outgoing | ",
			"response: ", resp,
			" error: ", err)

		return
	}
}
