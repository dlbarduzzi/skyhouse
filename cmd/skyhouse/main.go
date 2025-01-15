package main

import (
	"context"
	"os"

	"github.com/dlbarduzzi/skyhouse/internal/logging"
)

func main() {
	logger := logging.NewLoggerFromEnv().With("app", "skyhouse")

	ctx := context.Background()
	ctx = logging.LoggerWithContext(ctx, logger)

	if err := start(ctx); err != nil {
		logger.Error(err.Error())
		os.Exit(2)
	}
}

func start(ctx context.Context) error {
	logger := logging.LoggerFromContext(ctx)
	logger.Info("starting application")
	return nil
}
