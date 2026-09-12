package submissionsrepository

import (
	"github.com/jackc/pgx/v5"
	dbpostgres "github.com/markgredasov/viewcode-submissions/internal/database/postgres"
	errs "github.com/markgredasov/viewcode-submissions/internal/errors"
	"github.com/markgredasov/viewcode-submissions/internal/model/submission"
)

type Repository struct {
	pool dbpostgres.Pool
}

func New(pool dbpostgres.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func scanSubmission(row pgx.Row) (submission.Submission, error) {
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
		&rec.CreatedAt,
		&rec.UpdatedAt,
	)
	if err != nil {
		return submission.Submission{}, errs.Internal("scan submission", err)
	}

	return toModel(rec)
}

func scanSubmissions(rows pgx.Rows) ([]submission.Submission, error) {
	var out []submission.Submission

	for rows.Next() {
		s, err := scanSubmission(rows)
		if err != nil {
			return nil, errs.Internal("scan submissions", err)
		}
		out = append(out, s)
	}
	if err := rows.Err(); err != nil {
		return nil, errs.Internal("iterate submissions", err)
	}

	return out, nil
}
