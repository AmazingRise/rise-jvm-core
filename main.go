package main

import (
	"io/ioutil"
	"os"
	"rise-jvm-core/jvm"
	"rise-jvm-core/loader"
	"rise-jvm-core/logger"
)

func main() {
	//logger.InitLogger(os.Stdout, os.Stdout, os.Stdout)
	logger.InitLogger(ioutil.Discard, ioutil.Discard, os.Stdout)

	if len(os.Args) <= 1 {
		logger.Errorln("no input file")
	}
	file, _ := os.Open(os.Args[1])
	l := loader.CreateLoader()

	class := l.LoadClass(file)
	logger.Infoln("Class name: ", class.This, "@", class.Super)

	vm := jvm.CreateVM()
	vm.AppendClass(class)
	vm.Boot()
}
