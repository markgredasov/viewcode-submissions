package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/domain/submission"
	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
)

func (h *Handler) SetResult(ctx context.Context, req *submissions_v1.SetResultRequest) (*submissions_v1.SetResultResponse, error) {
	rh := response.NewHandler(logger.FromContext(ctx))

	status, err := submission.ConvertStatusFromGRPC(*req.GetStatus().Enum())
	if err != nil {
		return nil, rh.GRPCError(err, "convert status from grpc enum object")
	}

	updated, err := h.service.SetResult(ctx, req.GetId(), status, req.GetTestsPassed(), req.GetExecutionTimeMs())
	if err != nil {
		return nil, rh.GRPCError(err, "set result")
	}

	resp, err := setResultToResponse(updated)
	if err != nil {
		return nil, rh.GRPCError(err, "convert to response")
	}

	return resp, nil
}
