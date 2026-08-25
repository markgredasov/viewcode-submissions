package submissionsservice

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model"
)

type service struct {
	repository SubmissionsRepository
}

type SubmissionsRepository interface {
	Create(ctx context.Context, submission model.Submission) (model.Submission, error)
}

func New(repository SubmissionsRepository) *service {
	return &service{
		repository: repository,
	}
}
