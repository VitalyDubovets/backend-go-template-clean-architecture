package log

import (
	"backend/internal/infrastructure/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLog struct {
	logger *zap.SugaredLogger
}

func (z *ZapLog) Debug(args ...any) {
	z.logger.Debug(args)
}

func (z *ZapLog) Info(args ...any) {
	z.logger.Info(args)
}

func (z *ZapLog) Warn(args ...any) {
	z.logger.Warn(args)
}

func (z *ZapLog) Error(args ...any) {
	z.logger.Error(args)
}

func (z *ZapLog) Fatal(args ...any) {
	z.logger.Fatal(args)
}

func NewZapLog(config *config.AppConfig) (syncFn func(), err error) {
	logLvl, err := zapcore.ParseLevel(config.LogLevel)
	if err != nil {
		return nil, err
	}

	zapLog, err := zap.Config{
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
		return nil, err
	}

	zap.ReplaceGlobals(zapLog)

	log = &ZapLog{logger: zapLog.Sugar()}

	return nil, err
}
