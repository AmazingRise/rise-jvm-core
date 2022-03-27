package entity

type Attribute struct {
	Name  string
	Bytes []byte
}

func (a *Attribute) ConverToCode() (code ByteCode) {
	/*code.MaxStack = l.u2()            // max stack
	code.MaxLocals = l.u2()           // max locals
	codeLen := l.u4()                 // code length
	bytes = l.readBytes(int(codeLen)) // code
	exTableLen := l.u2()
	code.ExceptionTable = make([]entity.Exception, exTableLen)
	for i := 0; i < int(exTableLen); i++ {
		code.ExceptionTable[i].StartPc = l.u2()
		code.ExceptionTable[i].EndPc = l.u2()
		code.ExceptionTable[i].HandlerPc = l.u2()
		code.ExceptionTable[i].CatchType = l.u2()
	}
	//fmt.Println(code)
	l.readAttributes(l.u2())
	return*/
	return ByteCode{}
}

func (a *Attribute) readBytes(length int) []byte {
	return []byte{}
}
