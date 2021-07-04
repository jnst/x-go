package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// MustCreateLogger creates logger.
func CreateLogger() *zap.SugaredLogger {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)
	// zapcore
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
	logger, _ := config.Build()

	return logger.Sugar()
}
