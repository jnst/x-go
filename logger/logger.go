package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CreateSugaredLogger creates sugared logger.
func CreateSugaredLogger() *zap.SugaredLogger {
	logger, _ := createZapConfig()
	return logger.Sugar()
}

// CreateLogger creates logger.
func CreateLogger() *zap.Logger {
	logger, _ := createZapConfig()
	return logger
}

func createZapConfig() (*zap.Logger, error) {
	level := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config := zap.Config{
		Level:    level,
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	return config.Build()
}
