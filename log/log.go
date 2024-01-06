package log

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

func Trace(tpl string, args ...interface{}) {
	logImpl(log.Trace(), tpl, args...)
}

func Debug(tpl string, args ...interface{}) {
	logImpl(log.Debug(), tpl, args...)
}

func Info(tpl string, args ...interface{}) {
	logImpl(log.Info(), tpl, args...)
}

func Warn(tpl string, args ...interface{}) {
	logImpl(log.Warn(), tpl, args...)
}

func Error(tpl string, args ...interface{}) {
	logImpl(log.Error(), tpl, args...)
}

func Fatal(tpl string, args ...interface{}) {
	logImpl(log.Fatal(), tpl, args...)
}

func GetImpl() zerolog.Logger {
	return log.Logger
}

func logImpl(impl *zerolog.Event, tpl string, args ...interface{}) {
	msg := fmt.Sprintf(tpl, args...)
	impl.Msg(msg)
}
