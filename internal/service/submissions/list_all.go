package submissionsservice

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model"
)

func (s *service) ListAll(ctx context.Context) ([]model.Submission, error) {
	return s.repository.ListAll(ctx)
}
