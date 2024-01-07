package log

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var (
	logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: "2006-01-02T15:04:05.999",
	}).With().Timestamp().Logger()
)

func init() {
}

func SetLevel(level zerolog.Level) {
	logger = logger.Level(zerolog.DebugLevel)
}

func Trace(tpl string, args ...interface{}) {
	logImpl(logger.Trace(), tpl, args...)
}

func Debug(tpl string, args ...interface{}) {
	logImpl(logger.Debug(), tpl, args...)
}

func Info(tpl string, args ...interface{}) {
	logImpl(logger.Info(), tpl, args...)
}

func Warn(tpl string, args ...interface{}) {
	logImpl(logger.Warn(), tpl, args...)
}

func Error(tpl string, args ...interface{}) {
	logImpl(logger.Error(), tpl, args...)
}

func Panic(tpl string, args ...interface{}) {
	logImpl(logger.Panic(), tpl, args...)
}

func GetImpl() zerolog.Logger {
	return logger
}

func logImpl(impl *zerolog.Event, tpl string, args ...interface{}) {
	msg := fmt.Sprintf(tpl, args...)
	impl.Msg(msg)
}
