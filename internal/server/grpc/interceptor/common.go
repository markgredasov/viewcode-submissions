package interceptor

import (
	"context"
	"fmt"
	"time"

	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"github.com/markgredasov/viewcode-submissions/internal/model"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	loggerKey string = "logger"
)

func RequestID() Unary {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		const requestIDHeader = "x-request-id"

		requestID, err := getFromMetadata(ctx, requestIDHeader)
		if err != nil {
			return nil, errs.Internal("request id interceptor", err)
		}

		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(requestIDHeader, requestID))
		ctx = metadata.AppendToOutgoingContext(ctx, requestIDHeader, requestID)

		return handler(ctx, req)
	}
}

func Logger(l *logger.Logger) Unary {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		const requestIDHeader = "x-request-id"

		requestID, err := getFromMetadata(ctx, requestIDHeader)
		if err != nil {
			return nil, errs.Internal("logger interceptor", err)
		}

		logger := l.With(
			zap.String("Request-ID", requestID),
			zap.String("Method", info.FullMethod),
		)

		//nolint:staticcheck // not needed
		ctx = context.WithValue(ctx, loggerKey, logger)

		return handler(ctx, req)
	}
}

func Recovery() Unary {
	//nolint:nonamedreturns // more readable to use named returns here
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		log := logger.FromContext(ctx)
		defer func() {
			if r := recover(); r != nil {
				err = response.Panic(log, r, "during handle gRPC request got unexpected panic")
			}
		}()

		return handler(ctx, req)
	}
}

func Trace() Unary {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		log := logger.FromContext(ctx)
		start := time.Now().UTC()

		log.Debug(
			">>> incoming gRPC request",
			zap.Time("time", start),
		)

		resp, err := handler(ctx, req)

		statusCode := response.GetStatusCode(err)

		log.Debug(
			"<<< done gRPC request",
			zap.String("status-code", statusCode.String()),
			zap.Duration("latency", time.Since(start)),
		)

		return resp, err
	}
}

func getFromMetadata(ctx context.Context, header string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err := fmt.Errorf("failed to get metadata")
		return "", errs.Internal("get from metadata", err)
	}

	vals := md.Get(header)
	if len(vals) == 0 {
		vals = []string{model.NewUUID().String()}
	}
	val := vals[0]

	return val, nil
}
