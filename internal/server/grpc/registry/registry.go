package registry

import "google.golang.org/grpc"

type RegisterFunc func(*grpc.Server)

type Registry struct {
	registerFuncs []RegisterFunc
}

func New() *Registry {
	return &Registry{}
}

func (r *Registry) Add(registerFuncs ...RegisterFunc) *Registry {
	r.registerFuncs = append(r.registerFuncs, registerFuncs...)
	return r
}

func (r *Registry) RegisterAll(server *grpc.Server) {
	for _, registerFunc := range r.registerFuncs {
		registerFunc(server)
	}
}
