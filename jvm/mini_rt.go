package jvm

import "fmt"

// Mini Runtime

type Rt struct {
	MethodRef map[string]func(args ...interface{}) []interface{}
}

func CreateRt() *Rt {
	rt := &Rt{
		MethodRef: make(map[string]func(args ...interface{}) []interface{}),
	}
	rt.MethodRef["java/io/PrintStream.println"] = func(args ...interface{}) []interface{} {
		fmt.Println(args...)
		return nil
	}
	rt.MethodRef["java/io/PrintStream.print"] = func(args ...interface{}) []interface{} {
		fmt.Print(args...)
		return nil
	}
	return rt
}

func (r *Rt) Find(class string, method string) bool {
	_, ok := r.MethodRef[class+"."+method]
	return ok
}

func (r *Rt) Run(name string, args ...interface{}) []interface{} {
	fn := r.MethodRef[name]
	return fn(args...)
}
