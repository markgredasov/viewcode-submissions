package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/markgredasov/viewcode-submissions/internal/app/closer"
	"github.com/markgredasov/viewcode-submissions/internal/logger"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/interceptor"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/interceptor/iplimiter"
	grpcserver "github.com/markgredasov/viewcode-submissions/internal/server/grpc/server"
)

type App struct {
	server      *grpcserver.GRPCServer
	diContainer *diContainer
}

func New(ctx context.Context) *App {
	a := &App{
		diContainer: newDIContainer(),
	}

	a.initDeps(ctx)

	return a
}

func (a *App) initDeps(ctx context.Context) {
	inits := []func(context.Context){
		a.diContainer.initPool,
		a.initGRPCServer,
	}

	for _, fn := range inits {
		fn(ctx)
	}
}

func (a *App) Run(ctx context.Context) {
	if err := a.server.Run(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("run application:", err)
	}

	closerCtx, closerCancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Second) //nolint:mnd // todo
	defer closerCancel()

	if err := closer.CloseAll(closerCtx); err != nil {
		fmt.Println("failed to close deps:", err)
	}
}

func (a *App) initGRPCServer(ctx context.Context) {
	configIPLimiter := iplimiter.NewConfigMust()
	ipLimiter := iplimiter.New(configIPLimiter)

	log := logger.FromContext(ctx)

	configServer := grpcserver.NewConfigMust()
	server := grpcserver.New(
		configServer,
		log,
		interceptor.RequestID(),
		interceptor.Logger(log),
		interceptor.Recovery(),
		interceptor.RateLimiter(ipLimiter),
		interceptor.Trace(),
	)

	server.RegisterServices(a.diContainer.Methods().RegisterAll)

	a.server = server
}
