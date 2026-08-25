package app

import (
	server_grpc "github.com/markgredasov/viewcode-submissions/internal/server/grpc/server"
	"google.golang.org/grpc"
)

func (a *App) initGRPCServer(registrars []server_grpc.ServiceRegistrar) {
	serverConfig := server_grpc.NewConfigMust()
	server := server_grpc.NewGRPCServer(
		serverConfig,
		a.logger,
	)

	server.RegisterServices(func(s *grpc.Server) {
		server_grpc.RegisterServices(s, registrars...)
	})

	a.server = server
}
