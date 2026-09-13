package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/request"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
)

func (h *Handler) GetByID(ctx context.Context, req *submissions_v1.GetByIDRequest) (*submissions_v1.GetByIDResponse, error) {
	rh := response.NewHandler(logger.FromContext(ctx))

	if err := request.Validate(req); err != nil {
		return nil, rh.GRPCError(err, "validate request")
	}

	s, err := h.service.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, rh.GRPCError(err, "get submission by id")
	}

	return getByIDToResponse(s)
}
