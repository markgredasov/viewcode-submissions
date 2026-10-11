package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
)

func (h *Handler) SetRunning(ctx context.Context, req *submissions_v1.SetRunningRequest) (*submissions_v1.SetRunningResponse, error) {
	rh := response.NewHandler(logger.FromContext(ctx))

	updated, err := h.service.SetRunning(ctx, req.GetId())
	if err != nil {
		return nil, rh.GRPCError(err, "set running")
	}

	resp, err := setRunningToResponse(updated)
	if err != nil {
		return nil, rh.GRPCError(err, "convert to response")
	}

	return resp, nil
}
