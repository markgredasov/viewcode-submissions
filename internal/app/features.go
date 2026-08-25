package app

import (
	submissionshandler "github.com/markgredasov/viewcode-submissions/internal/handler/submissions"
	server_grpc "github.com/markgredasov/viewcode-submissions/internal/server/grpc/server"
)

func (a *App) initFeatures() ([]server_grpc.ServiceRegistrar, error) {
	registrars := []server_grpc.ServiceRegistrar{
		submissionshandler.NewRegistrar(nil),
	}

	return registrars, nil
}
