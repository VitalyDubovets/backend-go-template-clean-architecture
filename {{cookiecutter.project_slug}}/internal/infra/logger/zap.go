package logger

import (
	"{{ cookiecutter.project_slug }}/internal/infra/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLog struct {
	log *zap.SugaredLogger
}

func (z *ZapLog) Debug(args ...any) {
	z.log.Debug(args...)
}

func (z *ZapLog) Info(args ...any) {
	z.log.Info(args...)
}

func (z *ZapLog) Warn(args ...any) {
	z.log.Warn(args...)
}

func (z *ZapLog) Error(args ...any) {
	z.log.Error(args...)
}

func (z *ZapLog) Fatal(args ...any) {
	z.log.Fatal(args...)
}

func (z *ZapLog) Debugf(template string, args ...any) {
	z.log.Debugf(template, args...)
}

func (z *ZapLog) Infof(template string, args ...any) {
	z.log.Infof(template, args...)
}

func (z *ZapLog) Warnf(template string, args ...any) {
	z.log.Warnf(template, args...)
}

func (z *ZapLog) Errorf(template string, args ...any) {
	z.log.Errorf(template, args...)
}

func (z *ZapLog) Fatalf(template string, args ...any) {
	z.log.Fatalf(template, args...)
}

func NewZapLog(logConfig *config.LogConfig) (BaseLogger, func() error, error) {
	logLvl, err := zapcore.ParseLevel(logConfig.LogLevel)
	if err != nil {
		return nil, nil, err
	}

	zapLog, err := zap.Config{
		Encoding:         logConfig.LogFormat,
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

	syncFn := func() error {
		return zapLog.Sync()
	}

	return &ZapLog{log: zapLog.Sugar()}, syncFn, err
}
