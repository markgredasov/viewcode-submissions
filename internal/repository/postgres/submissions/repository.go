package submissionsrepository

import (
	"github.com/jackc/pgx/v5"
	dbpostgres "github.com/markgredasov/viewcode-submissions/internal/database/postgres"
	"github.com/markgredasov/viewcode-submissions/internal/model"
)

type repository struct {
	pool dbpostgres.Pool
}

func New(pool dbpostgres.Pool) *repository {
	return &repository{
		pool: pool,
	}
}

func scanSubmission(row pgx.Row) (model.Submission, error) {
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
