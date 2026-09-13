package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/request"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
)

func (h *Handler) Create(ctx context.Context, req *submissions_v1.CreateRequest) (*submissions_v1.CreateResponse, error) {
	rh := response.NewHandler(logger.FromContext(ctx))

	if err := request.Validate(req); err != nil {
		return nil, rh.GRPCError(err, "validate request")
	}

	requestModel, err := createToSubmission(req)
	if err != nil {
		return nil, rh.GRPCError(err, "map request to submission")
	}

	createdModel, err := h.service.Create(ctx, requestModel)
	if err != nil {
		return nil, rh.GRPCError(err, "create submission")
	}

	resp, err := createToResponse(createdModel)
	if err != nil {
		return nil, rh.GRPCError(err, "map submission to response")
	}

	return resp, nil
}
