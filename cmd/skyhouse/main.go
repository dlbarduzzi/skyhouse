package main

import (
	"context"
	"os"

	"github.com/dlbarduzzi/skyhouse/internal/logging"
	"github.com/dlbarduzzi/skyhouse/internal/server"
	"github.com/dlbarduzzi/skyhouse/internal/skyhouse"
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

	appConfig := setSkyhouseConfig()

	app, err := skyhouse.NewSkyhouse(appConfig, logger)
	if err != nil {
		return err
	}

	srv := server.NewServer(app.Port(), logger)

	srv.RunBeforeShutdown(func() {
		app.Shutdown()
	})

	return srv.Start(ctx, app.Routes())
}

// TODO: Get port number from config.
func setSkyhouseConfig() *skyhouse.Config {
	return &skyhouse.Config{
		Port: 8000,
	}
}
