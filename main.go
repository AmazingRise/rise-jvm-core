package main

import (
	"fmt"
	"os"
	"wasm-jvm/jvm"
	"wasm-jvm/loader"
	"wasm-jvm/utils"
)

func main() {
	logger := utils.CreateLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	file, _ := os.Open("./Add.class")
	l := loader.CreateLoader(logger)

	class := l.Load(file)
	fmt.Println("Class name: ", class.This, "@", class.Super)

	vm := jvm.CreateVM(logger)
	vm.AppendClass(class)
	//vm.Boot()
	vm.ExecStaticMethod("Add", "Add5")
}
