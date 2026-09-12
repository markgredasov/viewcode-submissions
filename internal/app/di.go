package app

import (
	"context"
	"os"

	"github.com/markgredasov/viewcode-submissions/internal/app/closer"
	dbpostgres "github.com/markgredasov/viewcode-submissions/internal/database/postgres"
	submissionshandler "github.com/markgredasov/viewcode-submissions/internal/handler/submissions"
	"github.com/markgredasov/viewcode-submissions/internal/logger"
	submissionsrepository "github.com/markgredasov/viewcode-submissions/internal/repository/postgres/submissions"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/registry"
	submissionsservice "github.com/markgredasov/viewcode-submissions/internal/service/submissions"
	"go.uber.org/zap"
)

type diContainer struct {
	pool *dbpostgres.ConnectionPool

	submissionsRepo    *submissionsrepository.Repository
	submissionsService *submissionsservice.Service
	submissionsHandler *submissionshandler.Handler

	methods *registry.Registry
}

func newDIContainer() *diContainer {
	d := &diContainer{}

	return d
}

func (d *diContainer) initPool(ctx context.Context) {
	log := logger.FromContext(ctx)
	config := dbpostgres.NewConfigMust()
	pool, err := dbpostgres.NewConnectionPool(ctx, config)
	if err != nil {
		log.Error("initialize pool", zap.Error(err))
		os.Exit(1)
	}

	closer.Add("connection pool", func(_ context.Context) error {
		pool.Close()
		return nil
	})

	d.pool = pool
}

func (d *diContainer) Pool() *dbpostgres.ConnectionPool {
	if d.pool == nil {
		panic("connection pool must be initialized with ctx before all initializations in container")
	}

	return d.pool
}

func (d *diContainer) SubmissionsRepository() *submissionsrepository.Repository {
	if d.submissionsRepo == nil {
		d.submissionsRepo = submissionsrepository.New(d.Pool())
	}

	return d.submissionsRepo
}

func (d *diContainer) SubmissionsService() *submissionsservice.Service {
	if d.submissionsService == nil {
		d.submissionsService = submissionsservice.New(d.SubmissionsRepository())
	}

	return d.submissionsService
}

func (d *diContainer) SubmissionsHandler() *submissionshandler.Handler {
	if d.submissionsHandler == nil {
		d.submissionsHandler = submissionshandler.New(d.SubmissionsService())
	}

	return d.submissionsHandler
}

func (d *diContainer) Methods() *registry.Registry {
	if d.methods == nil {
		d.methods = registry.New().Add(
			d.SubmissionsHandler().Register(),
		)
	}

	return d.methods
}
