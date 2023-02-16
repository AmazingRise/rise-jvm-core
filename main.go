package core

import (
	"github.com/AmazingRise/rise-jvm-core/jvm"
	"github.com/AmazingRise/rise-jvm-core/loader"
	"github.com/AmazingRise/rise-jvm-core/logger"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		panic("no input file")
	}
	RunFromPath(os.Args[1]+".class", os.Stdout, nil, true)
}

// Run WASM entrance
func Run(file io.Reader, out io.Writer, in io.Reader, silent bool, args ...string) {
	if silent {
		logger.InitLogger(ioutil.Discard, ioutil.Discard, os.Stdout)
	} else {
		logger.InitLogger(os.Stdout, os.Stdout, os.Stdout)
	}

	l := loader.CreateLoader()

	class := l.LoadClass(file)

	vm := jvm.CreateVM(out, in)
	vm.AppendClass(class)
	vm.Boot(args...)
}

func RunFromPath(path string, out io.Writer, in io.Reader, silent bool, args ...string) {
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
	vm.Boot(args...)
}
