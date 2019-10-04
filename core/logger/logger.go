package logger

import (
	"log"
	"os"
)

var (
	Log logger
)

type (
	Logger interface {
		Info(format string, v ...interface{})
		Error(format string, v ...interface{})
	}

	logger struct {
		err *log.Logger
		out *log.Logger
	}
)

func init() {
	Log = logger{
		err: log.New(os.Stderr, "", 0),
		out: log.New(os.Stdout, "", 0),
	}
}

func (l logger) Info(format string, v ...interface{}) {
	l.out.Printf(format, v...)
}

func (l logger) Error(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}

func Info(format string, v ...interface{}) {
	Log.Info(format, v...)
}

func Error(format string, v ...interface{}) {
	Log.Error(format, v...)
}
