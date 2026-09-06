package submissionsservice

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model"
	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
)

func (s *service) GetByID(ctx context.Context, id string) (submission.Submission, error) {
	uid, err := model.ToUUID(id)
	if err != nil {
		return submission.Submission{}, err
	}

	return s.repository.GetByID(ctx, uid)
}
