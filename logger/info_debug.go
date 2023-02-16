//go:build debug
// +build debug

package logger

func Infoln(v ...interface{}) {
	debugLogger.Println(v...)
}

func Infof(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}
