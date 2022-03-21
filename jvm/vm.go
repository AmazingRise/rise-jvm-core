package jvm

import (
	"fmt"
	"log"
	"wasm-jvm/loader"
	"wasm-jvm/utils"
)

type VM struct {
	threads []Thread
	classes map[string]*loader.Class
}

const StackSize = 100
const HeapSize = 1024

// Logger
var debugLog, infoLog, warnLog, errLog *log.Logger

func CreateVM(logger utils.Logger) *VM {
	errLog = logger.ErrorLogger
	warnLog = logger.WarnLogger
	infoLog = logger.InfoLogger
	debugLog = logger.DebugLogger

	vm := &VM{
		classes: make(map[string]*loader.Class),
	}

	return vm
}

func (v *VM) AppendClass(class *loader.Class) {
	v.classes[class.This] = class
}

// Boot to boot a JVM with loaded classes
func (v *VM) Boot() {
	// search for a public class with a static main method
	main := v.findMain()
	if main == nil {
		panic("classes does not contain a main")
	}
	v.Exec(main.Attrs.ByteCode)
}

func (v *VM) findMain() *loader.Method {
	for _, c := range v.classes {
		main, ok := c.Methods["main"]
		if c.IsPublic() && ok && main.IsPublic() && main.IsStatic() {
			return main
		}
	}
	return nil
}

func (v *VM) Exec(code loader.Code) {
	// byteCode := code.Bytes
	for _, b := range code.Bytes {
		fmt.Printf("%02X ", b)
	}
	//fmt.Println(code)
}

func (v *VM) ExecStaticMethod(className string, methodName string) {
	v.Exec(v.classes[className].Methods[methodName].Attrs.ByteCode)
}
