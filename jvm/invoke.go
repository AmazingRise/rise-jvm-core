package jvm

import (
	"strings"
	"wasm-jvm/entity"
	"wasm-jvm/logger"
)

// invoke create a frame
func (v *VM) invoke(method *entity.Method, args ...interface{}) *Frame {
	logger.Infoln("Calling method", method.Name, "with args:", args)
	// TODO: Some check, e.g. Abstract method should not be executed
	// Create frame
	frame := &Frame{}
	// If it is a patched method
	if method.This == nil {
		// This is a patched method
		frame.Stack = v.rt.Run(method.Name, args...)
		frame.State = FrameExit
		return frame
	}
	// Load Text
	for _, attr := range method.Attrs {
		// fmt.Printf("%02X ", attr.Text)
		if attr.Name == "Code" {
			frame.ByteCode = *Load(attr.Bytes)
			break
		}
	}
	// Load arguments
	frame.Locals = make([]interface{}, frame.MaxLocals)
	n := len(args)
	for i := n - 1; i >= 0; i-- {
		frame.Locals[i] = args[i]
	}
	frame.This = method.This
	frame.MethodName = method.Name

	return frame
}

func (v *VM) LocateMethod(className string, methodName string) *entity.Method {
	// TODO: Exception process
	// TODO: Overwrite
	// If it is runtime call
	if strings.HasPrefix(className, "java/") {
		if v.rt.Find(className, methodName) {
			return v.CreateDummyMethod(className, methodName)
		} else {
			logger.Errorf("unsupported runtime call: %s::%s", className, methodName)
		}
	}
	class, ok := v.classes[className]
	if !ok {
		logger.Errorln("unable to locate class", className)
	}
	method, ok := class.Methods[methodName]
	if !ok {
		logger.Errorf("unable to locate method %s in class %s.\n", className, methodName)
	}
	return method
}

func (v *VM) InvokeStaticMethod(method *entity.Method, args ...interface{}) *Frame {
	// TODO: Some check
	// TODO: Overwrite
	return v.invoke(method, args...)
}

func (v *VM) CreateDummyMethod(className string, methodName string) *entity.Method {
	return &entity.Method{
		Name:  className + "." + methodName,
		Flags: 0,
		Desc:  "",
		Attrs: []entity.Attribute{{
			Name:  "Code",
			Bytes: []byte{0xFF, 0xAC},
		}},
		This: nil,
	}
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
