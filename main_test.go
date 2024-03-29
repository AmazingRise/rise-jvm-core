package core

import (
	"bytes"
	"fmt"
	"github.com/AmazingRise/rise-jvm-core/jvm"
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
	buf := bytes.NewBufferString("")
	RunFromPath("demo/Benchmark.class", buf, nil, true)
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

func TestFor(t *testing.T) {
	Verify(t, "demo/For.class", "12345")
}

func TestMemoryLeak(t *testing.T) {
	Verify(t, "demo/MemoryLeak.class", "")
}

func TestOverload(t *testing.T) {
	Verify(t, "demo/Overload.class", "false\ntrue\nfalse\ntrue\n")
}

func TestPlus10(t *testing.T) {
	Verify(t, "demo/Plus10.class", "25\n", "15")
}

func TestFib(t *testing.T) {
	Verify(t, "/home/rise/Coding/rise-jvm-cli/Fib.class", "6765\n", "20")
}

func TestBenchmark(t *testing.T) {
	buf := bytes.NewBufferString("")
	RunFromPath("demo/Benchmark.class", buf, nil, true)
	fmt.Println(buf.String())
}

func TestAll(t *testing.T) {
	t.Run("Add", TestAdd)
	t.Run("Recursive", TestRecursive)
	t.Run("Object", TestObj)
	t.Run("For", TestFor)
	t.Run("Overload", TestOverload)
	//t.Run("Memory Leak", TestMemoryLeak)
}

func Verify(t *testing.T, path string, out string, args ...string) {
	buf := bytes.NewBufferString("")
	fmt.Println("Testing", path)
	RunFromPath(path, buf, nil, false, args...)
	if out != buf.String() {
		t.Errorf("excepted: %s\nbut: %s\n", out, buf.String())
	}
}

func TestDesc(t *testing.T) {
	fmt.Println(jvm.GetParamCount("(Ljava/lang/Object;)V"))
}
