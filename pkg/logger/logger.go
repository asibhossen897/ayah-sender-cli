package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(msg string, keysAndValues ...interface{}) {
	infoLogger.Println(append([]interface{}{msg}, keysAndValues...)...)
}

func Error(msg string, keysAndValues ...interface{}) {
	errorLogger.Println(append([]interface{}{msg}, keysAndValues...)...)
}
