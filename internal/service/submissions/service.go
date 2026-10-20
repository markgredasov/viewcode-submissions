package submissionsservice

import (
	"context"
	"uuid"

	"github.com/markgredasov/viewcode-submissions/internal/domain/submission"
)

type Service struct {
	repository SubmissionsRepository
	tx         TxManager
}

type SubmissionsRepository interface {
	Create(ctx context.Context, submission submission.Submission) (submission.Submission, error)
	GetByID(ctx context.Context, id uuid.UUID) (submission.Submission, error)
	ListAll(ctx context.Context) ([]submission.Submission, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status submission.Status, testsPassed int, executionTimeInMS int) (submission.Submission, error)
}

type TxManager interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}

func New(repository SubmissionsRepository, txManager TxManager) *Service {
	return &Service{
		repository: repository,
		tx:         txManager,
	}
}
