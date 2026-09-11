package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/request"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
)

func (h *Handler) GetByID(ctx context.Context, req *submissions_v1.GetByIDRequest) (*submissions_v1.GetByIDResponse, error) {
	if err := request.Validate(req); err != nil {
		return nil, response.GRPCError(err, "validate request")
	}

	s, err := h.service.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, response.GRPCError(err, "get submission by id")
	}

	return getByIDToResponse(s)
}
