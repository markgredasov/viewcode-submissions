package submissionsrepository

import database_postgres "github.com/markgredasov/viewcode-submissions/internal/database/postgres"

type repository struct {
	pool database_postgres.Pool
}

func New(pool database_postgres.Pool) *repository {
	return &repository{
		pool: pool,
	}
}
