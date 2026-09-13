package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) ListAll(ctx context.Context, _ *emptypb.Empty) (*submissions_v1.ListAllResponse, error) {
	rh := response.NewHandler(logger.FromContext(ctx))

	m, err := h.service.ListAll(ctx)
	if err != nil {
		return nil, rh.GRPCError(err, "list all submissions")
	}

	return listAllToResponse(m)
}
