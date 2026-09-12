package response

import (
	"errors"
	"fmt"

	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"go.uber.org/zap"
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
	case errors.Is(err, errs.ErrForbidden):
		statusCode = codes.PermissionDenied
	case errors.Is(err, errs.ErrTooManyRequests):
		statusCode = codes.ResourceExhausted
	default:
		statusCode = codes.Internal
	}

	return errorResponse(statusCode, err, msg)
}

func Panic(log *logger.Logger, p any, msg string) error {
	err := fmt.Errorf("unexpected panic: %v", p)
	log.Error(msg, zap.Error(err))

	return errorResponse(codes.Internal, err, msg)
}

func errorResponse(statusCode codes.Code, err error, msg string) error {
	return status.Error(statusCode, fmt.Sprintf("%s: %v", msg, err))
}

func GetStatusCode(err error) codes.Code {
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
