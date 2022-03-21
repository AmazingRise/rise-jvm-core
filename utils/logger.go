package utils

import (
	"fmt"
	"io"
	"log"
)

// logger has 4 types:
// DEBUG INFO WARN ERROR

type Logger struct {
	DebugLogger, InfoLogger, WarnLogger, ErrorLogger *log.Logger
}

func CreateLogger(debug, info, warn, error io.Writer) Logger {
	return Logger{
		DebugLogger: log.New(debug, "[DEBUG] ", log.LstdFlags),
		InfoLogger:  log.New(debug, "[INFO] ", log.LstdFlags),
		WarnLogger:  log.New(debug, "[WARN] ", log.LstdFlags),
		ErrorLogger: log.New(debug, "[ERROR] ", log.LstdFlags),
	}
}

func printHex(bytes []byte) {
	for _, b := range bytes {
		fmt.Printf("%x", b)
	}
}
