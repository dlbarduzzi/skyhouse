package main

import (
	"github.com/dlbarduzzi/skyhouse/internal/logging"
)

func main() {
	logger := logging.NewLoggerFromEnv().With("app", "skyhouse")
	logger.Info("running application")
}
