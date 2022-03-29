package jvm

import (
	"wasm-jvm/entity"
	"wasm-jvm/logger"
)

// invoke create a frame
func (v *VM) invoke(method *entity.Method, args ...interface{}) *Frame {
	logger.Infoln("Calling method", method.Name, "with args:", args)
	// TODO: Some check, e.g. Abstract method should not be executed
	// Create frame
	frame := &Frame{}
	// Load Text
	for _, attr := range method.Attrs {
		// fmt.Printf("%02X ", attr.Text)
		if attr.Name == "Code" {
			frame.ByteCode = *Load(attr.Bytes)
			break
		}
	}
	//logger.Infoln(thread.Text)
	// Load arguments
	frame.Locals = make([]interface{}, frame.MaxLocals)
	n := len(args)
	for i := n - 1; i >= 0; i-- {
		frame.Locals[i] = args[i]
	}
	frame.This = method.This

	// Append to threads
	//v. = append(v.frame, frame)
	return frame
}

func (v *VM) LocateMethod(className string, methodName string) *entity.Method {
	// TODO: Exception process
	// TODO: Overwrite
	return v.classes[className].Methods[methodName]
}

func (v *VM) InvokeStaticMethod(method *entity.Method, args ...interface{}) *Frame {
	// TODO: Some check
	// TODO: Overwrite
	return v.invoke(method, args...)
}

// bootstrap find main and put the frame into a new thread
func (v *VM) bootstrap() {
	// search for a public class with a static main method
	main := v.findMain()
	if main == nil {
		logger.Errorln("classes does not contain a main")
	}
	frame := v.invoke(main)
	v.pool.CreateThread(frame)
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
