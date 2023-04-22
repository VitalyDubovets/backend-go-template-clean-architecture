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
	Debugf(template string, args ...any)
	Infof(template string, args ...any)
	Warnf(template string, args ...any)
	Errorf(template string, args ...any)
	Fatalf(template string, args ...any)
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

func (b *BaseLog) Debugf(template string, args ...any) {
	b.log.Printf(template, args)
}

func (b *BaseLog) Infof(template string, args ...any) {
	b.log.Printf(template, args)
}

func (b *BaseLog) Warnf(template string, args ...any) {
	b.log.Printf(template, args)
}

func (b *BaseLog) Errorf(template string, args ...any) {
	b.log.Printf(template, args)
}

func (b *BaseLog) Fatalf(template string, args ...any) {
	b.log.Fatalf(template, args)
}

func GetLogInstance() BaseLogger {
	if log == nil {
		once.Do(func() {
			log = &BaseLog{log: l.Default()}
		})
	}

	return log
}
