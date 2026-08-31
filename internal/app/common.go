package app

import (
	dbpostgres "github.com/markgredasov/viewcode-submissions/internal/database/postgres"
	"github.com/markgredasov/viewcode-submissions/internal/logger"
)

func (a *App) initLogger() error {
	loggerConfig := logger.NewConfigMust()
	logger, err := logger.NewLogger(loggerConfig)
	if err != nil {
		return err
	}
	a.logger = logger

	return nil
}

func (a *App) initPostgresPool() error {
	databaseConfig := dbpostgres.NewConfigMust()
	postgresPool, err := dbpostgres.NewConnectionPool(a.ctx, databaseConfig)
	if err != nil {
		return err
	}

	a.pool = postgresPool

	return nil
}
