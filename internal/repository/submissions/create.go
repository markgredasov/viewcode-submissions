package submissionsrepository

import (
	"context"

	"github.com/markgredasov/viewcode-submissions/internal/model"
)

func (r *repository) Create(ctx context.Context, submission model.Submission) (model.Submission, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	sqlQuery := `
		INSERT INTO users.submissions (id, user_id, problem_id, tests_total, tests_passed, condition, lang, execution_time_ms)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, user_id, problem_id, tests_total, tests_passed, condition, lang, execution_time_ms;
	`

	row := r.pool.QueryRow(
		ctx,
		sqlQuery,
		submission.ID,
		submission.UserID,
		submission.ProblemID,
		submission.TestsTotal,
		submission.TestsPassed,
		submission.Status,
		submission.Language,
		submission.ExecutionTimeInMS,
	)

	var rec record
	err := row.Scan(
		&rec.ID,
		&rec.UserID,
		&rec.ProblemID,
		&rec.TestsTotal,
		&rec.TestsPassed,
		&rec.Status,
		&rec.Language,
		&rec.ExecutionTimeInMS,
	)
	if err != nil {
		return model.Submission{}, err
	}

	return fromRecordToModel(rec)
}
