package runtime

import (
	"fmt"
	"io"
	"rise-jvm-core/entity"
	"rise-jvm-core/logger"
)

// Mini Runtime

type Rt struct {
	MethodRef map[string]*Method
	Object    *entity.Class
}

type Method struct {
	Flags uint16
	Fn    func(args ...interface{}) []interface{}
}

func CreateRt(out io.Writer, in io.Reader) *Rt {
	rt := &Rt{
		MethodRef: make(map[string]*Method),
	}
	rt.MethodRef["java/io/PrintStream.println"] = &Method{
		Flags: entity.MethodFinal,
		Fn: func(args ...interface{}) []interface{} {
			_, _ = fmt.Fprintln(out, args[1:]...)
			return nil
		},
	}

	rt.MethodRef["java/io/PrintStream.print"] = &Method{
		Flags: entity.MethodFinal,
		Fn: func(args ...interface{}) []interface{} {
			_, _ = fmt.Fprint(out, args[1:]...)
			return nil
		},
	}

	rt.MethodRef["java/lang/Object.<init>"] = &Method{
		Flags: 0,
		Fn: func(args ...interface{}) []interface{} {
			return nil
		},
	}

	rt.MethodRef["java/lang/Boolean.valueOf"] = &Method{
		Flags: entity.MethodStatic,
		Fn: func(args ...interface{}) []interface{} {
			return []interface{}{args[0].(int) == 1}
		},
	}

	rt.Object = &entity.Class{
		Name:      "java/lang/Object",
		Methods:   nil,
		Constants: nil,
		Super:     "",
		Flags:     entity.ClassPublic,
	}
	return rt
}

func (r *Rt) LocateMethod(class string, method string, desc string) *entity.Method {
	rtMethod, ok := r.MethodRef[class+"."+method]
	if !ok {
		return nil
	}
	return &entity.Method{
		Name:  method,
		Flags: rtMethod.Flags,
		Desc:  desc,
		Attrs: nil,
		This:  &entity.Class{Name: class},
	}
}

func (r *Rt) RunMethod(name string, args ...interface{}) []interface{} {
	logger.Infoln("Runtime method", name, "executed with", args)
	fn := r.MethodRef[name].Fn
	return fn(args...)
}

func (r *Rt) CreateFakeClass(name string) *entity.Class {
	return &entity.Class{
		Name:      name,
		Methods:   nil,
		Constants: nil,
		Super:     "",
		Flags:     0,
	}
}

// CreateDummyMethod deprecated to avoid memory leaking
func (r *Rt) CreateDummyMethod(className string, methodName string) *entity.Method {
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
