package app

import (
	submissionshandler "github.com/markgredasov/viewcode-submissions/internal/handler/submissions"
	submissionsrepository "github.com/markgredasov/viewcode-submissions/internal/repository/submissions"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/registry"
	submissionsservice "github.com/markgredasov/viewcode-submissions/internal/service/submissions"
)

func (a *App) initFeatures() *registry.Registry {
	reg := registry.New()

	reg.Add(
		a.initFeatureSubmissions(),
	)

	return reg
}

func (a *App) initFeatureSubmissions() registry.RegisterFunc {
	repository := submissionsrepository.New(a.pool)
	service := submissionsservice.New(repository)
	handler := submissionshandler.New(service)

	return handler.Register()
}
