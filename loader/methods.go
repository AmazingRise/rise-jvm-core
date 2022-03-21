package loader

/*
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type Method struct {
	Name  string
	Flags uint16
	Attrs Attributes
}

func (m *Method) IsPublic() bool {
	return m.Flags&0x0001 > 0
}

func (m *Method) IsStatic() bool {
	return m.Flags&0x0008 > 0
}

func (l *ClassLoader) readMethods(count uint16) {
	var i uint16
	c := l.class
	c.Methods = make(map[string]*Method)
	for i = 0; i < count; i++ {
		method := &Method{}
		method.Flags = l.u2() // access flags
		debugLog.Printf("Method flags: %b", method.Flags)
		nameIdx := l.u2() // name index
		method.Name = c.Constants.getUtf8Constant(nameIdx)
		debugLog.Println("Method Name: ", method.Name)
		l.u2() // descriptor index
		aCount := l.u2()
		method.Attrs = l.readAttributes(aCount)
		c.Methods[method.Name] = method
	}
}
