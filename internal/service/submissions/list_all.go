package submissionsservice

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
)

func (s *service) ListAll(ctx context.Context) ([]submission.Submission, error) {
	return s.repository.ListAll(ctx)
}
