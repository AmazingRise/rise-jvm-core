package main

import (
	"io/ioutil"
	"os"
	"wasm-jvm/jvm"
	"wasm-jvm/loader"
	"wasm-jvm/logger"
)

func main() {
	//logger.InitLogger(os.Stdout, os.Stdout, os.Stdout)
	logger.InitLogger(ioutil.Discard, ioutil.Discard, ioutil.Discard)

	file, _ := os.Open("./Add.class")
	l := loader.CreateLoader()

	class := l.LoadClass(file)
	logger.Infoln("Class name: ", class.This, "@", class.Super)

	vm := jvm.CreateVM()
	vm.AppendClass(class)
	vm.Boot()
}
