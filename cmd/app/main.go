package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/markgredasov/viewcode-submissions/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	app.New(ctx).Run(ctx)
}
