package log

import (
	"backend/internal/infrastructure/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(config *config.AppConfig) (*zap.Logger, error) {
	logLvl, err := zapcore.ParseLevel(config.LogLevel)
	if err != nil {
		return nil, err
	}

	logger, err := zap.Config{
		Encoding:         config.LogFormat,
		Level:            zap.NewAtomicLevelAt(logLvl),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "timestamp",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,

			FunctionKey: "func",
		},
	}.Build()

	return logger, err
}
