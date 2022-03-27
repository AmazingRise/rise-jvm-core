package jvm

import (
	"fmt"
	"wasm-jvm/entity"
)

type VM struct {
	threads []Thread
	classes map[string]*entity.Class
}

const StackSize = 100
const HeapSize = 1024

func CreateVM() *VM {
	vm := &VM{
		classes: make(map[string]*entity.Class),
	}

	return vm
}

func (v *VM) AppendClass(class *entity.Class) {
	v.classes[class.This] = class
}

// Boot to boot a JVM with loaded classes
func (v *VM) Boot() {
	// search for a public class with a static main method
	main := v.findMain()
	if main == nil {
		panic("classes does not contain a main")
	}
	v.Exec(main.Code)
}

func (v *VM) findMain() *entity.Method {
	for _, c := range v.classes {
		main, ok := c.Methods["main"]
		if c.IsPublic() && ok && main.IsPublic() && main.IsStatic() {
			return main
		}
	}
	return nil
}

func (v *VM) Exec(code entity.ByteCode) {
	// byteCode := code.Bytes
	for _, b := range code.Bytes {
		fmt.Printf("%02X ", b)
	}
	//fmt.Println(code)
}

func (v *VM) ExecStaticMethod(className string, methodName string) {
	v.Exec(v.classes[className].Methods[methodName].Code)
}
