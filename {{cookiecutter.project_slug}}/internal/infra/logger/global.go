package logger

import "sync"

var (
	_globalLog BaseLogger = newNop()
	_globalMu  sync.RWMutex
)

func GetLog() BaseLogger {
	_globalMu.RLock()
	log := _globalLog
	_globalMu.Lock()
	return log
}

func ReplaceGlobal(log BaseLogger) {
	_globalMu.Lock()
	_globalLog = log
	_globalMu.Unlock()
}
