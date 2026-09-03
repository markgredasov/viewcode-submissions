package submissionsrepository

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model"
)

func (r *repository) GetByID(ctx context.Context, id model.UUID) (model.Submission, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	sqlQuery := `
		SELECT id, user_id, problem_id, tests_total, tests_passed, condition, lang, execution_time_ms
		FROM users.submissions
		WHERE id = $1;
	`

	row := r.pool.GetQuerier(ctx).QueryRow(ctx, sqlQuery, id)

	return scanSubmission(row)
}
