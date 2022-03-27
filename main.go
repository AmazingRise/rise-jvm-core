package main

import (
	"os"
	"wasm-jvm/jvm"
	"wasm-jvm/loader"
	"wasm-jvm/logger"
)

func main() {
	logger.InitLogger(os.Stdout, os.Stdout, os.Stdout)

	file, _ := os.Open("./Add.class")
	l := loader.CreateLoader()

	class := l.LoadClass(file)
	logger.Infoln("Class name: ", class.This, "@", class.Super)

	vm := jvm.CreateVM()
	vm.AppendClass(class)
	//vm.Boot()
	add := vm.LocateMethod("Add", "Add")
	vm.InvokeStaticMethod(add, 1, 2)
	vm.Schedule()
}
