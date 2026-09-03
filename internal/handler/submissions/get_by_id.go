package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/request"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
)

func (h *handler) GetByID(ctx context.Context, req *submissions_v1.GetByIDRequest) (*submissions_v1.GetByIDResponse, error) {
	if err := request.Validate(req); err != nil {
		return nil, response.GRPCError(err, err.Error())
	}

	s, err := h.service.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, response.GRPCError(err, err.Error())
	}

	return getByIDToResponse(s)
}
