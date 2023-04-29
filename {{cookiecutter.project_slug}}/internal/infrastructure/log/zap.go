package log

import (
	"backend/internal/infrastructure/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLog struct {
	log *zap.SugaredLogger
}

func (z *zapLog) Debug(args ...any) {
	z.log.Debug(args)
}

func (z *zapLog) Info(args ...any) {
	z.log.Info(args)
}

func (z *zapLog) Warn(args ...any) {
	z.log.Warn(args)
}

func (z *zapLog) Error(args ...any) {
	z.log.Error(args)
}

func (z *zapLog) Fatal(args ...any) {
	z.log.Fatal(args)
}

func (z *zapLog) Debugf(template string, args ...any) {
	z.log.Debugf(template, args)
}

func (z *zapLog) Infof(template string, args ...any) {
	z.log.Infof(template, args)
}

func (z *zapLog) Warnf(template string, args ...any) {
	z.Warnf(template, args)
}

func (z *zapLog) Errorf(template string, args ...any) {
	z.Errorf(template, args)
}

func (z *zapLog) Fatalf(template string, args ...any) {
	z.Fatalf(template, args)
}

func NewZapLog(config *config.LogConfig) (BaseLogger, func(), error) {
	logLvl, err := zapcore.ParseLevel(config.LogLevel)
	if err != nil {
		return nil, nil, err
	}

	log, err := zap.Config{
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

	if err != nil {
		return nil, nil, err
	}

	return &zapLog{log: log.Sugar()}, nil, err
}
