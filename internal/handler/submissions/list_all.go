package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/response"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handler) ListAll(ctx context.Context, _ *emptypb.Empty) (*submissions_v1.ListAllResponse, error) {
	m, err := h.service.ListAll(ctx)
	if err != nil {
		return nil, response.GRPCError(err, "list all submissions")
	}

	return listAllToResponse(m)
}
