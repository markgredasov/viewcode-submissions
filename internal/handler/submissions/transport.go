package submissionshandler

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/registry"
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
	"google.golang.org/grpc"
)

type handler struct {
	service SubmissionsService
}

type SubmissionsService interface {
	Create(ctx context.Context, submission submission.Submission) (submission.Submission, error)
	GetByID(ctx context.Context, id string) (submission.Submission, error)
	ListAll(ctx context.Context) ([]submission.Submission, error)
}

func New(service SubmissionsService) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Register() registry.RegisterFunc {
	return func(s *grpc.Server) {
		submissions_v1.RegisterSubmissionsServiceServer(s, h)
	}
}
