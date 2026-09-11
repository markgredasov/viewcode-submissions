package submissionsservice

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model"
	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
)

type Service struct {
	repository SubmissionsRepository
}

type SubmissionsRepository interface {
	Create(ctx context.Context, submission submission.Submission) (submission.Submission, error)
	GetByID(ctx context.Context, id model.UUID) (submission.Submission, error)
	ListAll(ctx context.Context) ([]submission.Submission, error)
}

func New(repository SubmissionsRepository) *Service {
	return &Service{
		repository: repository,
	}
}
