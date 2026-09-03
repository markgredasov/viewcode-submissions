package app

import (
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/interceptor"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/registry"
	grpcserver "github.com/markgredasov/viewcode-submissions/internal/server/grpc/server"
)

func (a *App) initGRPCServer(reg *registry.Registry) {
	serverConfig := grpcserver.NewConfigMust()
	server := grpcserver.New(
		serverConfig,
		a.logger,
		interceptor.RequestID(),
		interceptor.Logger(a.logger),
		interceptor.Recovery(),
		interceptor.Trace(),
	)

	server.RegisterServices(reg.RegisterAll)

	a.server = server
}
