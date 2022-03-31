package main

import (
	"bytes"
	"os"
	"runtime/pprof"
	"testing"
)

type Case struct {
	Path   string
	Output string
}

func TestMemory(t *testing.T) {
	f, _ := os.OpenFile("mem.profile", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()

	for i := 0; i <= 1000; i++ {
		main()
	}
	pprof.Lookup("heap").WriteTo(f, 0)
}

func TestAdd(t *testing.T) {
	Verify(t, "demo/Add.class", "Hello world!\n(100+200)*300=90000\n")
}

func TestRecursive(t *testing.T) {
	Verify(t, "demo/Recursive.class", "Hello world!\n(100+200)*300=90000\n")
}

func TestObj(t *testing.T) {
	Verify(t, "demo/Obj.class", "1\n")
}

func TestAll(t *testing.T) {
	t.Run("Add", TestAdd)
	t.Run("Recursive", TestRecursive)
	t.Run("Object", TestObj)
}

func Verify(t *testing.T, path string, out string) {
	buf := bytes.NewBufferString("")
	LoadAndRun(path, buf, nil, false)
	if out != buf.String() {
		t.Errorf("excepted: %s\nbut: %s\n", out, buf.String())
	}
}
