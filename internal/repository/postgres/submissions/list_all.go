package submissionsrepository

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
)

func (r *Repository) ListAll(ctx context.Context) ([]submission.Submission, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	sqlQuery := `
		SELECT id, user_id, problem_id, tests_total, tests_passed, condition, lang, execution_time_ms, created_at, updated_at
		FROM users.submissions;
	`

	rows, err := r.pool.GetQuerier(ctx).Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanSubmissions(rows)
}
