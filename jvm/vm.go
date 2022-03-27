package jvm

import (
	"wasm-jvm/entity"
	"wasm-jvm/logger"
)

type VM struct {
	threads []*Thread
	classes map[string]*entity.Class
}

func CreateVM() *VM {
	vm := &VM{
		classes: make(map[string]*entity.Class),
	}

	return vm
}

func (v *VM) AppendClass(class *entity.Class) {
	v.classes[class.This] = class
}

// Boot to boot a JVM
func (v *VM) Boot() {
	v.InvokeMain()
	v.Schedule()
}

// Schedule find ready threads and execute it, then switch between threads
func (v *VM) Schedule() {
	// Currently, we only support single thread
	for _, thread := range v.threads {
		if thread.State == ThreadReady {
			result := thread.Exec()
			logger.Infoln("Thread exits, with a result", result)
		}
	}
}
