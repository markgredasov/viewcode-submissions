package submissionsservice

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/domain/submission"
)

func (s *Service) Create(ctx context.Context, submission submission.Submission) (submission.Submission, error) {
	return s.repository.Create(ctx, submission)
}
