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

type handler struct {
	log *logger.Logger
}

func NewHandler(log *logger.Logger) *handler {
	return &handler{
		log: log,
	}
}

func (h *handler) GRPCError(err error, msg string) error {
	var (
		statusCode codes.Code
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, errs.ErrAlreadyExists):
		statusCode = codes.AlreadyExists
		logFunc = h.log.Error
	case errors.Is(err, errs.ErrInvalidArgument):
		statusCode = codes.InvalidArgument
		logFunc = h.log.Info
	case errors.Is(err, errs.ErrNotFound):
		statusCode = codes.NotFound
		logFunc = h.log.Info
	case errors.Is(err, errs.ErrNotImplemented):
		statusCode = codes.Unimplemented
		logFunc = h.log.Error
	case errors.Is(err, errs.ErrForbidden):
		statusCode = codes.PermissionDenied
		logFunc = h.log.Info
	case errors.Is(err, errs.ErrTooManyRequests):
		statusCode = codes.ResourceExhausted
		logFunc = h.log.Info
	default:
		statusCode = codes.Internal
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))

	return h.errorResponse(statusCode, err, msg)
}

func (h *handler) Panic(p any, msg string) error {
	err := fmt.Errorf("unexpected panic: %v", p)
	h.log.Error(msg, zap.Error(err))

	return h.errorResponse(codes.Internal, err, msg)
}

func (h *handler) errorResponse(statusCode codes.Code, err error, msg string) error {
	return status.Error(statusCode, fmt.Sprintf("%s: %v", msg, err))
}
