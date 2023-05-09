package logger

import (
	l "log"
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

type baseLog struct {
	log *l.Logger
}

func (b *baseLog) Debug(args ...any) {
	b.log.Println(args...)
}

func (b *baseLog) Info(args ...any) {
	b.log.Println(args...)
}

func (b *baseLog) Warn(args ...any) {
	b.log.Println(args...)
}

func (b *baseLog) Error(args ...any) {
	b.log.Println(args...)
}

func (b *baseLog) Fatal(args ...any) {
	b.log.Fatalln(args...)
}

func (b *baseLog) Debugf(template string, args ...any) {
	b.log.Printf(template, args...)
}

func (b *baseLog) Infof(template string, args ...any) {
	b.log.Printf(template, args...)
}

func (b *baseLog) Warnf(template string, args ...any) {
	b.log.Printf(template, args...)
}

func (b *baseLog) Errorf(template string, args ...any) {
	b.log.Printf(template, args...)
}

func (b *baseLog) Fatalf(template string, args ...any) {
	b.log.Fatalf(template, args...)
}

func newNop() *baseLog {
	return &baseLog{log: l.Default()}
}
