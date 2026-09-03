package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type Unary func(context.Context, any, *grpc.UnaryServerInfo, grpc.UnaryHandler) (any, error)

func Chain(interceptor ...Unary) grpc.ServerOption {
	var out []grpc.UnaryServerInterceptor
	for _, i := range interceptor {
		out = append(out, grpc.UnaryServerInterceptor(i))
	}

	return grpc.ChainUnaryInterceptor(out...)
}
