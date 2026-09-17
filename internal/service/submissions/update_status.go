package submissionsservice

import (
	"context"

	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
	"github.com/markgredasov/viewcode-submissions/internal/model"
	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
)

func (s *Service) SetRunning(ctx context.Context, id string) (submission.Submission, error) {
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

		if err = subm.AdvanceTo(submission.StatusRunning); err != nil {
			return err
		}

		updated, err = s.repository.UpdateStatus(ctx, uid, subm.Status, subm.TestsPassed, subm.ExecutionTimeInMS)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return submission.Submission{}, err
	}

	return updated, nil
}

func (s *Service) SetResult(ctx context.Context, id string, next submission.Status, testsPassed int32, executionTimeInMS int32) (submission.Submission, error) {
	uid, err := model.ToUUID(id)
	if err != nil {
		return submission.Submission{}, err
	}

	if !next.IsTerminal() {
		return submission.Submission{}, errs.Conflict("status is not terminal")
	}

	var updated submission.Submission
	err = s.tx.WithTx(ctx, func(ctx context.Context) error {
		var subm submission.Submission
		subm, err = s.repository.GetByID(ctx, uid)
		if err != nil {
			return err
		}

		if err = subm.AdvanceTo(next); err != nil {
			return err
		}

		updated, err = s.repository.UpdateStatus(ctx, uid, next, int(testsPassed), int(executionTimeInMS))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return submission.Submission{}, err
	}

	return updated, nil
}
