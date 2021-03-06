package loader

import (
	"github.com/AmazingRise/rise-jvm-core/entity"
)

/*
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

func (l *Loader) readMethods(count uint16) {
	var i uint16
	c := l.class
	c.Methods = make(map[string][]*entity.Method)
	for i = 0; i < count; i++ {
		method := &entity.Method{}
		method.Flags = l.u2() // access flags
		//logger.Infof("Method flags: %b", method.Flags)
		nameIdx := l.u2() // name index
		method.Name = c.Constants.GetUtf8Const(nameIdx)
		//logger.Infoln("Method Name: ", method.Name)
		descIdx := l.u2() // descriptor index
		method.Desc = c.Constants.GetUtf8Const(descIdx)
		//logger.Infoln("Method descriptor:", method.Desc)
		aCount := l.u2() // attribute count
		method.Attrs = l.ReadAttributes(aCount)
		method.This = c
		_, ok := c.Methods[method.Name]
		if ok {
			c.Methods[method.Name] = append(c.Methods[method.Name], method)
		} else {
			c.Methods[method.Name] = []*entity.Method{method}
		}
	}
}

/*
	Every Java Virtual Machine implementation must recognize Code attributes. If the
	method is either native or abstract, its method_info structure must not have a
	Code attribute. Otherwise, its method_info structure must have exactly one Code
	attribute.
*/
