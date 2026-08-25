package response

import (
	"errors"

	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GRPCError(err error, msg string) error {
	var statusCode codes.Code

	switch {
	case errors.Is(err, errs.ErrAlreadyExists):
		statusCode = codes.AlreadyExists
	case errors.Is(err, errs.ErrInvalidArgument):
		statusCode = codes.InvalidArgument
	case errors.Is(err, errs.ErrNotFound):
		statusCode = codes.NotFound
	case errors.Is(err, errs.ErrNotImplemented):
		statusCode = codes.Unimplemented
	default:
		statusCode = codes.Internal
	}

	return status.Error(statusCode, msg)
}
