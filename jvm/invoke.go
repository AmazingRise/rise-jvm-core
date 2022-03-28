package jvm

import (
	"wasm-jvm/entity"
	"wasm-jvm/logger"
)

// invoke a method, create a thread, and put it into v.threads
// execution is done by Scheduler
func (v *VM) invoke(method *entity.Method, args ...interface{}) {
	logger.Infoln("Calling method", method.Name, "with args:", args)
	// TODO: Some check, e.g. Abstract method should not be executed
	// Create thread
	thread := &Thread{}
	// Load Text
	for _, attr := range method.Attrs {
		// fmt.Printf("%02X ", attr.Text)
		if attr.Name == "Code" {
			thread.ByteCode = *Load(attr.Bytes)
			break
		}
	}
	//logger.Infoln(thread.Text)
	// Load arguments
	n := len(args)
	for i := n - 1; i >= 0; i-- {
		thread.Locals = append(thread.Locals, args[i])
	}

	// Append to threads
	v.threads = append(v.threads, thread)
}

func (v *VM) LocateMethod(className string, methodName string) *entity.Method {
	// TODO: Exception process
	return v.classes[className].Methods[methodName]
}

func (v *VM) InvokeStaticMethod(method *entity.Method, args ...interface{}) {
	// TODO: Some check
	v.invoke(method, args...)
}

func (v *VM) InvokeMain() {
	// search for a public class with a static main method
	main := v.findMain()
	if main == nil {
		logger.Errorln("classes does not contain a main")
	}
	v.invoke(main)
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
