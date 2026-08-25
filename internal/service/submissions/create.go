package submissionsservice

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model"
)

func (s *service) Create(ctx context.Context, submission model.Submission) (model.Submission, error) {
	return s.repository.Create(ctx, submission)
}
