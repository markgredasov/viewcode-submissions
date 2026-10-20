package submissionsservice

import (
	"context"
	"uuid"

	"github.com/markgredasov/viewcode-submissions/internal/domain/submission"
)

func (s *Service) GetByID(ctx context.Context, id string) (submission.Submission, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return submission.Submission{}, err
	}

	return s.repository.GetByID(ctx, uid)
}
