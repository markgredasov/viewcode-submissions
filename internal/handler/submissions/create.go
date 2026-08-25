package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/request"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
)

func (h *handler) Create(ctx context.Context, req *submissions_v1.CreateRequest) (*submissions_v1.CreateResponse, error) {
	if err := request.Validate(req); err != nil {
		return nil, response.GRPCError(err, err.Error())
	}

	requestModel, err := fromRequestToModel(req)
	if err != nil {
		return nil, response.GRPCError(err, err.Error())
	}

	createdModel, err := h.service.Create(ctx, requestModel)
	if err != nil {
		return nil, response.GRPCError(err, err.Error())
	}

	resp, err := fromModelToResponse(createdModel)
	if err != nil {
		return nil, response.GRPCError(err, err.Error())
	}

	return resp, nil
}
