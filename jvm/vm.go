package jvm

import (
	"wasm-jvm/entity"
)

type VM struct {
	pool    *ThreadPool
	classes map[string]*entity.Class
	rt      *Rt
}

func CreateVM() *VM {
	vm := &VM{
		classes: make(map[string]*entity.Class),
		rt:      CreateRt(),
	}
	vm.pool = vm.CreateThreadPool()

	return vm
}

func (v *VM) AppendClass(class *entity.Class) {
	v.classes[class.This] = class
}

// Boot to boot a JVM
func (v *VM) Boot() {
	v.bootstrap()
	v.pool.Schedule()
}
