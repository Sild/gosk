package log

import (
	"testing"
)

func TestExamples(t *testing.T) {
	Trace("trace msg")
	Debug("debug msg")
	Info("info msg")
	Warn("warn msg")
	Error("error msg")

	defer func() {
		if err := recover(); err != nil {
			Info("panic occurred")
		}
	}()
	Panic("panic msg")
}
