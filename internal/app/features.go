package app

import (
	submissionshandler "github.com/markgredasov/viewcode-submissions/internal/handler/submissions"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/registry"
)

func (a *App) initFeatures() *registry.Registry {
	reg := registry.New()

	reg.Add(
		a.initFeatureSubmissions(),
	)

	return reg
}

func (a *App) initFeatureSubmissions() registry.RegisterFunc {
	handler := submissionshandler.New(nil)

	return handler.Register()
}
