package jvm

import (
	"io"
	"rise-jvm-core/entity"
	"rise-jvm-core/runtime"
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
func (v *VM) Boot() {
	v.bootstrap()
	v.pool.Schedule()
}
