package interceptor

import (
	"context"

	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
}

func Validate() Unary {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		rh := response.NewHandler(logger.FromContext(ctx))

		if v, ok := req.(validator); ok {
			if err := v.Validate(); err != nil {
				err = errs.InvalidArgument(err.Error())
				return nil, rh.GRPCError(err, "validate request")
			}
		}

		return handler(ctx, req)
	}
}
