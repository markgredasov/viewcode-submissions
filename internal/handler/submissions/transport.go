package submissionshandler

import (
	submissions_v1 "github.com/markgredasov/viewcode-submissions/pkg/api/v1"
	"google.golang.org/grpc"
)

type handler struct {
	service SubmissionsService
}

type SubmissionsService interface{}

func New(service SubmissionsService) *handler {
	return &handler{
		service: service,
	}
}

type Registrar struct {
	service *handler
}

func NewRegistrar(service *handler) *Registrar {
	return &Registrar{
		service: service,
	}
}

func (r *Registrar) RegisterGRPCService(server *grpc.Server) {
	submissions_v1.RegisterSubmissionsServiceServer(server, r.service)
}
