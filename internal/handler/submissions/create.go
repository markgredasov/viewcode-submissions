package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/request"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
)

func (h *handler) Create(ctx context.Context, req *submissions_v1.CreateRequest) (*submissions_v1.CreateResponse, error) {
	if err := request.Validate(req); err != nil {
		return nil, response.GRPCError(err, "validate request")
	}

	requestModel, err := createToSubmission(req)
	if err != nil {
		return nil, response.GRPCError(err, "map request to submission")
	}

	createdModel, err := h.service.Create(ctx, requestModel)
	if err != nil {
		return nil, response.GRPCError(err, "create submission")
	}

	resp, err := createToResponse(createdModel)
	if err != nil {
		return nil, response.GRPCError(err, "map submission to response")
	}

	return resp, nil
}
