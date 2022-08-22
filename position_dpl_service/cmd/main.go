package main

import (
	"github.com/404th/go_grpc_project/position_dpl_service/config"
	"github.com/404th/go_grpc_project/position_dpl_service/pkg/logger"
)

func main() {
	// loading config
	cfg := config.Load()

	// setting logger
	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	// getting postgres db

	// setting grpc server

	// listening port

	// serving to listening port
}
