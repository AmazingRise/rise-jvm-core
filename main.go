package main

import (
	"fmt"
	"os"
	"wasm-jvm/jvm"
	"wasm-jvm/loader"
	"wasm-jvm/logger"
)

func main() {
	logger.InitLogger(os.Stdout, os.Stdout, os.Stdout)

	file, _ := os.Open("./Add.class")
	l := loader.CreateLoader()

	class := l.Load(file)
	fmt.Println("Class name: ", class.This, "@", class.Super)

	vm := jvm.CreateVM()
	vm.AppendClass(class)
	//vm.Boot()
	vm.ExecStaticMethod("Add", "Add5")
}
