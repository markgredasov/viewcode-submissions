package server_grpc

import "google.golang.org/grpc"

type ServiceRegistrar interface {
	RegisterGRPCService(server *grpc.Server)
}

func RegisterServices(server *grpc.Server, registrars ...ServiceRegistrar) {
	for _, registrar := range registrars {
		registrar.RegisterGRPCService(server)
	}
}
