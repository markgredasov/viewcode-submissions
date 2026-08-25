package server_grpc

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	server *grpc.Server
	config Config
	log    *logger.Logger
}

func NewGRPCServer(config Config, logger *logger.Logger) *GRPCServer {
	return &GRPCServer{
		server: grpc.NewServer(),
		config: config,
		log:    logger,
	}
}

func (g *GRPCServer) RegisterServices(registerFunc func(*grpc.Server)) {
	registerFunc(g.server)

	if g.config.EnableReflection {
		reflection.Register(g.server)
	}
}

func (g *GRPCServer) Run(ctx context.Context) error {
	listenConfig := net.ListenConfig{}

	listener, err := listenConfig.Listen(ctx, "tcp", g.config.Address)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	defer listener.Close()

	ch := make(chan error, 1)

	go func() {
		defer func() {
			close(ch)
		}()

		g.log.Warn(
			"gRPC server started",
			zap.String("address", g.config.Address),
		)

		err = g.server.Serve(listener)

		if !errors.Is(err, grpc.ErrServerStopped) {
			ch <- err
		}
	}()

	select {
	case err = <-ch:
		return fmt.Errorf("serve gRPC server: %w", err)
	case <-ctx.Done():
		g.log.Warn("shutdown gRPC server")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			g.config.ShutdownTimeout,
		)
		defer cancel()

		done := make(chan struct{})
		go func() {
			g.server.GracefulStop()
			close(done)
		}()

		select {
		case <-done:
			g.log.Warn("gRPC server stopped")
		case <-shutdownCtx.Done():
			g.server.Stop()
			return fmt.Errorf("shutdown gRPC server forcefully")
		}
	}

	return nil
}
