package submissionsrepository

import (
	"context"

	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
	"github.com/markgredasov/viewcode-submissions/internal/model"
	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
)

func (r *Repository) UpdateStatus(ctx context.Context, id model.UUID, status submission.Status) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	q := r.pool.GetQuerier(ctx)

	sqlQuery := `
		UPDATE users.submissions
		SET condition = $1,
			updated_at = NOW()
		WHERE id = $2;
	`

	if _, err := q.Exec(ctx, sqlQuery, status.String(), id.String()); err != nil {
		return errs.Internal("update submission status", err)
	}

	return nil
}
