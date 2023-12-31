package logger

import (
	"github.com/blazee5/auth-microservice/internal/config"
	"go.uber.org/zap"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Logger struct {
	Log *zap.SugaredLogger
}

func NewLogger(cfg *config.Config) *Logger {
	var log *zap.Logger

	switch cfg.Env {
	case envLocal:
		log, _ = zap.NewDevelopment()
	case envDev:
		log, _ = zap.NewDevelopment()
	case envProd:
		log, _ = zap.NewProduction()
	}

	defer log.Sync()

	return &Logger{Log: log.Sugar()}
}
