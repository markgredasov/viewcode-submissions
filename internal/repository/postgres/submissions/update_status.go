package submissionsrepository

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model"
	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
)

func (r *Repository) UpdateStatus(ctx context.Context, id model.UUID, status submission.Status,
	testsPassed int, executionTimeInMS int) (submission.Submission, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	q := r.pool.GetQuerier(ctx)

	sqlQuery := `
		UPDATE users.submissions
		SET condition = $1,
			tests_passed = $2,
			execution_time_ms = $3,
			updated_at = NOW()
		WHERE id = $4
		RETURNING id, user_id, problem_id, tests_total, tests_passed, condition, lang, execution_time_ms, created_at, updated_at;
	`

	row := q.QueryRow(ctx, sqlQuery, status.String(), testsPassed, executionTimeInMS, id.String())

	return scanSubmission(row)
}
