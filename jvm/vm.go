package jvm

import (
	"github.com/AmazingRise/rise-jvm-core/entity"
	"github.com/AmazingRise/rise-jvm-core/logger"
	"github.com/AmazingRise/rise-jvm-core/runtime"
	"io"
)

type VM struct {
	pool    *ThreadPool
	classes map[string]*entity.Class
	rt      *runtime.Rt
}

func CreateVM(out io.Writer, in io.Reader) *VM {
	vm := &VM{
		classes: make(map[string]*entity.Class),
		rt:      runtime.CreateRt(out, in),
	}
	vm.pool = vm.CreateThreadPool()

	return vm
}

func (v *VM) AppendClass(class *entity.Class) {
	v.classes[class.Name] = class
}

// Boot to boot a JVM
func (v *VM) Boot(args ...string) {
	// search for a public class with a static main method
	main := v.findMain()
	if main == nil {
		logger.Errorln("classes does not contain a main")
	}

	arr := make([]interface{}, len(args))
	for i := range arr {
		arr[i] = args[i]
	}

	frame := v.InvokeMethod(main, arr)
	v.pool.CreateThread(frame)
	v.pool.Schedule()
}
