package interceptor

import (
	"context"
	"fmt"
	"time"

	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"github.com/markgredasov/viewcode-submissions/internal/model"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		log := logger.FromContext(ctx)
		var err error
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("during handle gRPC request got unexpected panic: %v", r)
				msg := "unexpected panic"
				err = errs.Internal(msg, err)
				log.Error(msg, zap.Error(err))
			}
		}()

		resp, hErr := handler(ctx, req)
		if hErr != nil {
			err = hErr
		}

		return resp, err
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

		statusCode := getStatusCode(err)

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

func getStatusCode(err error) codes.Code {
	statusCode := codes.OK
	if err != nil {
		if st, ok := status.FromError(err); ok {
			statusCode = st.Code()
		} else {
			statusCode = codes.Unknown
		}
	}
	return statusCode
}
