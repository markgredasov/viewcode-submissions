package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/markgredasov/viewcode-submissions/internal/app"
	"github.com/markgredasov/viewcode-submissions/internal/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logConfig := logger.NewConfigMust()
	log, err := logger.NewLogger(logConfig)
	if err != nil {
		fmt.Println("failed to initialize logger:", err)
		return
	}

	ctx = log.InContext(ctx)

	app.New(ctx).Run(ctx)
}
