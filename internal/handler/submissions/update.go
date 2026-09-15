package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/request"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
)

func (h *Handler) UpdateStatus(ctx context.Context, req *submissions_v1.UpdateStatusRequest) (*submissions_v1.UpdateStatusResponse, error) {
	rh := response.NewHandler(logger.FromContext(ctx))

	if err := request.Validate(req); err != nil {
		return nil, rh.GRPCError(err, "validate request")
	}

	status, err := submission.ConvertStatusFromGRPC(req.GetStatus())
	if err != nil {
		return nil, rh.GRPCError(err, "convert status from grpc")
	}

	updated, err := h.service.UpdateStatus(ctx, req.GetId(), status)
	if err != nil {
		return nil, rh.GRPCError(err, "update status")
	}

	resp, err := updateStatusToResponse(updated)
	if err != nil {
		return nil, rh.GRPCError(err, "convert to response")
	}

	return resp, nil
}
