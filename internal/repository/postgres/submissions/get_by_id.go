package submissionsrepository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
	"github.com/markgredasov/viewcode-submissions/internal/model"
	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
)

func (r *repository) GetByID(ctx context.Context, id model.UUID) (submission.Submission, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	sqlQuery := `
		SELECT id, user_id, problem_id, tests_total, tests_passed, condition, lang, execution_time_ms
		FROM users.submissions
		WHERE id = $1;
	`

	row := r.pool.GetQuerier(ctx).QueryRow(ctx, sqlQuery, id)

	s, err := scanSubmission(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return submission.Submission{}, errs.NotFound(fmt.Sprintf("no submission with id=%s found", id))
		}
		return submission.Submission{}, err
	}
	return s, nil
}
