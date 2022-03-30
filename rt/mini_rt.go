package rt

import (
	"fmt"
	"io"
)

// Mini Runtime

type Rt struct {
	MethodRef map[string]func(args ...interface{}) []interface{}
}

func CreateRt(out io.Writer, in io.Reader) *Rt {
	rt := &Rt{
		MethodRef: make(map[string]func(args ...interface{}) []interface{}),
	}
	rt.MethodRef["java/io/PrintStream.println"] = func(args ...interface{}) []interface{} {
		_, _ = fmt.Fprintln(out, args...)
		return nil
	}
	rt.MethodRef["java/io/PrintStream.print"] = func(args ...interface{}) []interface{} {
		_, _ = fmt.Fprint(out, args...)
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
