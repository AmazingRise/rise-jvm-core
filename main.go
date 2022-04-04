package main

import (
	"io"
	"io/ioutil"
	"os"
	"rise-jvm-core/jvm"
	"rise-jvm-core/loader"
	"rise-jvm-core/logger"
)

func main() {
	if len(os.Args) <= 1 {
		panic("no input file")
	}
	RunFromPath(os.Args[1]+".class", os.Stdout, nil, true)
}

// Run WASM entrance
func Run(file io.Reader, out io.Writer, in io.Reader, silent bool) {
	if silent {
		logger.InitLogger(ioutil.Discard, ioutil.Discard, os.Stdout)
	} else {
		logger.InitLogger(os.Stdout, os.Stdout, os.Stdout)
	}

	l := loader.CreateLoader()

	class := l.LoadClass(file)

	vm := jvm.CreateVM(out, in)
	vm.AppendClass(class)
	vm.Boot()
}

func RunFromPath(path string, out io.Writer, in io.Reader, silent bool) {
	if silent {
		logger.InitLogger(ioutil.Discard, ioutil.Discard, os.Stdout)
	} else {
		logger.InitLogger(os.Stdout, os.Stdout, os.Stdout)
	}

	file, _ := os.Open(path)
	l := loader.CreateLoader()

	class := l.LoadClass(file)

	vm := jvm.CreateVM(out, in)
	vm.AppendClass(class)
	vm.Boot()
}
