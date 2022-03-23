package loader

import "wasm-jvm/entity"

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

func (l *ClassLoader) readAttributes(count uint16) []entity.Attribute {
	var result []entity.Attribute
	var i uint16
	for i = 0; i < count; i++ {
		nameIdx := l.u2()
		name := l.class.Constants.GetUtf8Constant(nameIdx)
		debugLog.Println("Attribute:", name)
		aLen := l.u4() // attribute length
		bytes := l.readBytes(int(aLen))
		result = append(result, entity.Attribute{name, bytes})
	}
	return result
}

/*
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
	return
}
*/
