package app

import (
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/registry"
	server_grpc "github.com/markgredasov/viewcode-submissions/internal/server/grpc/server"
)

func (a *App) initGRPCServer(reg *registry.Registry) {
	serverConfig := server_grpc.NewConfigMust()
	server := server_grpc.NewGRPCServer(
		serverConfig,
		a.logger,
	)

	server.RegisterServices(reg.RegisterAll)

	a.server = server
}
