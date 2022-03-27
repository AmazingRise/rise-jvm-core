package entity

type ByteCode struct {
	MaxStack       uint16
	MaxLocals      uint16
	Text           []byte
	ExceptionTable []Exception
}

type Exception struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}
