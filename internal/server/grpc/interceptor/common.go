package interceptor

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/interceptor/iplimiter"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

func RequestID() Unary {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		const requestIDHeader = "x-request-id"

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.MD{}
		}

		requestID := getValFromMetadata(md, requestIDHeader)
		if requestID == "" {
			requestID = uuid.NewString()
		}

		md.Set(requestIDHeader, requestID)

		ctx = metadata.NewIncomingContext(ctx, md)
		ctx = metadata.AppendToOutgoingContext(ctx, requestIDHeader, requestID)

		return handler(ctx, req)
	}
}

func Logger(l *logger.Logger) Unary {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		const requestIDHeader = "x-request-id"

		requestID := metadata.ValueFromIncomingContext(ctx, requestIDHeader)[0]

		log := l.With(
			zap.String("Request-ID", requestID),
			zap.String("Method", info.FullMethod),
		)

		ctx = log.InContext(ctx)

		return handler(ctx, req)
	}
}

func Recovery() Unary {
	//nolint:nonamedreturns // more readable to use named returns here
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		h := response.NewHandler(logger.FromContext(ctx))

		defer func() {
			if r := recover(); r != nil {
				err = h.Panic(r, "during handle gRPC request got unexpected panic")
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

func RateLimiter(il *iplimiter.IPLimiter) Unary {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		h := response.NewHandler(logger.FromContext(ctx))

		p, ok := peer.FromContext(ctx)
		if !ok {
			err := fmt.Errorf("no peer found")
			return nil, h.GRPCError(errs.Internal("no peer found", err), "no peer found")
		}

		ip, _, err := net.SplitHostPort(p.Addr.String())
		if err != nil {
			ip = p.Addr.String()
		}

		if !il.Get(ip).Allow() {
			return nil, h.GRPCError(errs.TooManyRequests(), "rate limiter")
		}

		return handler(ctx, req)
	}
}

func getValFromMetadata(md metadata.MD, header string) string {
	vals := md.Get(header)
	if len(vals) == 0 {
		return ""
	}
	val := vals[0]

	return val
}
