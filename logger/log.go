package logger

import (
	"fmt"
	"io"
	"log"
)

var debugLogger, warnLogger, errLogger *log.Logger

func InitLogger(debug, warn, error io.Writer) {
	debugLogger = log.New(debug, "[DEBUG] ", log.LstdFlags)
	warnLogger = log.New(warn, "[WARN] ", log.LstdFlags)
	errLogger = log.New(error, "[ERROR] ", log.LstdFlags)
}

func Warnln(v ...interface{}) {
	warnLogger.Println(v...)
}

func Warnf(format string, v ...interface{}) {
	warnLogger.Printf(format, v...)
}

func Errorln(v ...interface{}) {
	errLogger.Fatalln(v...)
}

func Errorf(format string, v ...interface{}) {
	errLogger.Fatalf(format, v...)
}

func PrintHex(bytes []byte) {
	for _, b := range bytes {
		fmt.Printf("%x", b)
	}
}
