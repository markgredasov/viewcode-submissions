package submissionsrepository

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model"
)

func (r *repository) ListAll(ctx context.Context) ([]model.Submission, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	sqlQuery := `
		SELECT id, user_id, problem_id, tests_total, tests_passed, condition, lang, execution_time_ms
		FROM users.submissions;
	`

	rows, err := r.pool.GetQuerier(ctx).Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	return scanSubmissions(rows)
}
