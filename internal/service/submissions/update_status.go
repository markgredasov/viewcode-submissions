package submissionsservice

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model"
	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
)

func (s *Service) UpdateStatus(ctx context.Context, id string, nextStatus submission.Status) (submission.Submission, error) {
	uid, err := model.ToUUID(id)
	if err != nil {
		return submission.Submission{}, err
	}

	var updated submission.Submission
	err = s.tx.WithTx(ctx, func(ctx context.Context) error {
		var subm submission.Submission
		subm, err = s.repository.GetByID(ctx, uid)
		if err != nil {
			return err
		}

		if err = subm.AdvanceTo(nextStatus); err != nil {
			return err
		}

		if err = s.repository.UpdateStatus(ctx, uid, subm.Status); err != nil {
			return err
		}

		if updated, err = s.repository.GetByID(ctx, uid); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return submission.Submission{}, err
	}

	return updated, nil
}
