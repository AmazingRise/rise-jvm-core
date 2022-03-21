package loader

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

type Attributes struct {
	ByteCode Code
}

func (l *ClassLoader) readAttributes(count uint16) (result Attributes) {
	var i uint16
	for i = 0; i < count; i++ {
		nameIdx := l.u2()
		name := l.class.Constants.getUtf8Constant(nameIdx)
		debugLog.Println("Attribute:", name)
		aLen := l.u4() // attribute length
		//l.readBytes(int(aLen))
		switch name {
		case "Code":
			result.ByteCode = l.readCode()
		default:
			debugLog.Println("Unknown attribute, skipped")
			l.readBytes(int(aLen))
		}
	}
	return
}

/*
Code_attribute {
    u2 attribute_name_index; // read
    u4 attribute_length; // read
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type Code struct {
	MaxStack       uint16
	MaxLocals      uint16
	Bytes          []byte
	ExceptionTable []Exception
}

type Exception struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func (l *ClassLoader) readCode() (code Code) {
	code.MaxStack = l.u2()                 // max stack
	code.MaxLocals = l.u2()                // max locals
	codeLen := l.u4()                      // code length
	code.Bytes = l.readBytes(int(codeLen)) // code
	exTableLen := l.u2()
	code.ExceptionTable = make([]Exception, exTableLen)
	for i := 0; i < int(exTableLen); i++ {
		code.ExceptionTable[i].StartPc = l.u2()
		code.ExceptionTable[i].EndPc = l.u2()
		code.ExceptionTable[i].HandlerPc = l.u2()
		code.ExceptionTable[i].CatchType = l.u2()
	}
	//fmt.Println(code)
	l.readAttributes(l.u2())
	//code.Attributes =
	return
}

func (l *ClassLoader) readLineNumberTable() {

}

/*
	Every Java Virtual Machine implementation must recognize Code attributes. If the
	method is either native or abstract, its method_info structure must not have a
	Code attribute. Otherwise, its method_info structure must have exactly one Code
	attribute.
*/
