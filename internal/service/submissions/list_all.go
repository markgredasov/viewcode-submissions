package submissionsservice

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/domain/submission"
)

func (s *Service) ListAll(ctx context.Context) ([]submission.Submission, error) {
	return s.repository.ListAll(ctx)
}
