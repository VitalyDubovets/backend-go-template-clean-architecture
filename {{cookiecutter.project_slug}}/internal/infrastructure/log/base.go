package log

import (
	l "log"
	"sync"
)

var (
	log  BaseLogger
	once sync.Once
)

type BaseLogger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}

type BaseLog struct {
	log *l.Logger
}

func (b *BaseLog) Debug(args ...any) {
	b.log.Println(args)
}

func (b *BaseLog) Info(args ...any) {
	b.log.Println(args)
}

func (b *BaseLog) Warn(args ...any) {
	b.log.Println(args)
}

func (b *BaseLog) Error(args ...any) {
	b.log.Println(args)
}

func (b *BaseLog) Fatal(args ...any) {
	b.log.Fatalln(args)
}

func GetLogInstance() BaseLogger {
	if log == nil {
		once.Do(func() {
			log = &BaseLog{log: l.Default()}
		})
	}

	return log
}
