package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(logLevel string, env string) error {
	config := zap.NewProductionConfig()
	if env == "dev" {
		config = zap.NewDevelopmentConfig()
	}

	var lvl zapcore.Level
	switch logLevel {
	case "debug":
		lvl = zap.DebugLevel
	case "info":
		lvl = zap.InfoLevel
	case "warn":
		lvl = zap.WarnLevel
	case "error":
		lvl = zap.ErrorLevel
	}
	config.Level = zap.NewAtomicLevelAt(lvl)

	var err error
	logger, err := config.Build()
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	return nil
}
