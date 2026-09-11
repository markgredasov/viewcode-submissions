package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/markgredasov/viewcode-submissions/internal/app/closer"
	"github.com/markgredasov/viewcode-submissions/internal/server/grpc/interceptor"
	grpcserver "github.com/markgredasov/viewcode-submissions/internal/server/grpc/server"
)

type App struct {
	server      *grpcserver.GRPCServer
	diContainer *diContainer
}

func New(ctx context.Context) *App {
	a := &App{
		diContainer: newDIContainer(ctx),
	}

	a.initDeps()

	return a
}

func (a *App) initDeps() {
	inits := []func(){
		a.initGRPCServer,
	}

	for _, fn := range inits {
		fn()
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

func (a *App) initGRPCServer() {
	config := grpcserver.NewConfigMust()
	server := grpcserver.New(
		config,
		a.diContainer.log,
		interceptor.RequestID(),
		interceptor.Logger(a.diContainer.Logger()),
		interceptor.Recovery(),
		interceptor.Trace(),
	)

	server.RegisterServices(a.diContainer.Methods().RegisterAll)

	a.server = server
}
