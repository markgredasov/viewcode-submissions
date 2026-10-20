package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/domain/submission"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/registry"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
	"google.golang.org/grpc"
)

type Handler struct {
	service SubmissionsService
}

type SubmissionsService interface {
	Create(ctx context.Context, submission submission.Submission) (submission.Submission, error)
	GetByID(ctx context.Context, id string) (submission.Submission, error)
	ListAll(ctx context.Context) ([]submission.Submission, error)
	SetRunning(ctx context.Context, id string) (submission.Submission, error)
	SetResult(ctx context.Context, id string, status submission.Status, testsPassed int32, executionTimeInMS int32) (submission.Submission, error)
}

func New(service SubmissionsService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register() registry.RegisterFunc {
	return func(s *grpc.Server) {
		submissions_v1.RegisterSubmissionsServiceServer(s, h)
	}
}
