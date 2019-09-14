package logger

import "log"

func Info(format string, v ...interface{}) {
	log.Printf(format, v)
}
