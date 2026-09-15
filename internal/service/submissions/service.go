package submissionsservice

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model"
	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
)

type Service struct {
	repository SubmissionsRepository
	tx         TxManager
}

type SubmissionsRepository interface {
	Create(ctx context.Context, submission submission.Submission) (submission.Submission, error)
	GetByID(ctx context.Context, id model.UUID) (submission.Submission, error)
	ListAll(ctx context.Context) ([]submission.Submission, error)
	UpdateStatus(ctx context.Context, id model.UUID, status submission.Status) error
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
