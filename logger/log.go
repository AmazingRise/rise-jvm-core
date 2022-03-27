package logger

import (
	"fmt"
	"io"
	"log"
)

var debugLogger, warnLogger, errLogger *log.Logger

func InitLogger(debug, warn, error io.Writer) {
	debugLogger = log.New(debug, "[INFO] ", log.LstdFlags)
	warnLogger = log.New(debug, "[WARN] ", log.LstdFlags)
	errLogger = log.New(debug, "[ERROR] ", log.LstdFlags)
}

func Infoln(v ...interface{}) {
	debugLogger.Println(v)
}

func Infof(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}

func Warnln(v ...interface{}) {
	warnLogger.Println(v)
}

func Warnf(format string, v ...interface{}) {
	warnLogger.Printf(format, v...)
}

func Errorln(v ...interface{}) {
	errLogger.Fatalln(v)
}

func Errorf(format string, v ...interface{}) {
	errLogger.Fatalf(format, v...)
}

func printHex(bytes []byte) {
	for _, b := range bytes {
		fmt.Printf("%x", b)
	}
}
