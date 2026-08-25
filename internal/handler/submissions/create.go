package submissionshandler

import (
	"context"

	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) Create(ctx context.Context, req *submissions_v1.CreateRequest) (*submissions_v1.CreateResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %w", err)
	}
	return nil, nil
}
