package jvm

import (
	"bytes"
	"github.com/AmazingRise/rise-jvm-core/entity"
	"github.com/AmazingRise/rise-jvm-core/logger"
	"github.com/AmazingRise/rise-jvm-core/utils"
	"strings"
)

// invoke create a frame
func (v *VM) invoke(method *entity.Method, args ...interface{}) *Frame {
	var frame *Frame
	// Load Text
	for _, attr := range method.Attrs {
		// fmt.Printf("%02X ", attr.Text)
		if attr.Name == "Code" {
			byteCode := *Load(attr.Bytes)
			frame = CreateFrame(int(byteCode.MaxStack))
			frame.ByteCode = byteCode
			break
		}
	}

	logger.Infoln("Calling method", method.Name, "with args:", args)
	// TODO: Some check, e.g. Abstract method should not be executed

	// Load arguments
	frame.Locals = make([]interface{}, frame.MaxLocals)
	n := len(args)
	for i := n - 1; i >= 0; i-- {
		frame.Locals[i] = args[i]
	}
	frame.This = method.This
	frame.MethodRef = method

	return frame
}

func (v *VM) LocateMethod(className string, methodName string, desc string) *entity.Method {
	// TODO: Exception process
	// TODO: Overwrite
	// If it is runtime call
	if strings.HasPrefix(className, "java/") {
		rtMtd := v.rt.LocateMethod(className, methodName, desc)
		if rtMtd != nil {
			return rtMtd
		} else {
			logger.Errorf("unsupported runtime call: %s::%s", className, methodName)
		}
	}
	class, ok := v.classes[className]
	if !ok {
		logger.Errorln("unable to locate class", className)
	}
	methods, ok := class.Methods[methodName]
	if !ok {
		logger.Errorf("unable to locate method %s in class %s.\n", className, methodName)
	}
	for _, method := range methods {
		if method.Desc == desc {
			return method
		}
	}
	logger.Errorf("unable to locate method %s with description %s in class %s.\n", className, methodName, desc)
	return nil
}

func (v *VM) LocateClass(className string) *entity.Class {
	if className == "java/lang/Object" {
		return v.rt.Object
	}
	class, ok := v.classes[className]
	if !ok {
		if strings.HasPrefix(className, "java/") {
			fake := v.rt.CreateFakeClass(className)
			v.AppendClass(fake)
		} else {
			logger.Errorln("unable to locate class", className)
		}
	}
	return class
}

func (v *VM) InvokeMethod(method *entity.Method, args ...interface{}) *Frame {
	// TODO: Some check
	return v.invoke(method, args...)
}

func (v *VM) InvokeRuntimeMethod(method *entity.Method, args ...interface{}) *Frame {
	frame := CreateFrame(0)
	frame.Text = []byte{OpReturn}
	logger.Infoln("Calling runtime method "+method.Name+" with args ", args)
	//frame.Stack = v.rt.RunMethod(method.This.Name+"."+method.Name, args...)
	result := v.rt.RunMethod(method.This.Name+"."+method.Name, args...)
	if len(result) != 0 {
		frame.DataStack.Push(result[0])
		frame.Text = []byte{OpAReturn}
	}
	//logger.Infoln("Runtime method "+method.Name+" returns", frame.Stack)
	//frame.State = FrameExit
	return frame
}

func (v *VM) InvokeVirtualMethod(method *entity.Method, args ...interface{}) *Frame {
	frame := v.invoke(method, args...)
	return frame
}

func (v *VM) findMain() *entity.Method {
	for _, c := range v.classes {
		if !c.IsPublic() {
			continue
		}
		main, ok := c.Methods["main"]
		if ok && len(main) == 1 && main[0].IsPublic() && main[0].IsStatic() {
			return main[0]
		}
	}
	return nil
}

func Load(raw []byte) *entity.ByteCode {
	code := &entity.ByteCode{}
	r := utils.CreateReader(bytes.NewReader(raw))
	code.MaxStack = r.U2()
	code.MaxLocals = r.U2()
	codeLen := r.U4()
	code.Text = r.ReadBytes(int(codeLen))
	exLen := r.U2()
	var i uint16
	for i = 0; i < exLen; i++ {
		code.ExceptionTable = append(code.ExceptionTable, entity.Exception{
			StartPc:   r.U2(),
			EndPc:     r.U2(),
			HandlerPc: r.U2(),
			CatchType: r.U2(),
		})
	}
	// TODO: Read attributes
	return code
}
