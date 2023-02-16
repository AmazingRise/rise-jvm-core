package jvm

import (
	"github.com/AmazingRise/rise-jvm-core/entity"
)

type Frame struct {
	entity.ByteCode
	PC uint32
	//Stack  []interface{}
	DataStack *ArrStack
	Locals    []interface{}
	//Ins    [][]byte

	MethodRef *entity.Method
	State     int

	This *entity.Class
}

const (
	FrameReady = 0
	FrameExit  = 1
	FramePush  = 2
)

func CreateFrame(maxStack int) *Frame {
	return &Frame{DataStack: CreateArrStack(maxStack)}
}
func (f *Frame) condJmp(cond bool) {
	if cond {
		offset := int16(f.Text[f.PC+1])<<8 + int16(f.Text[f.PC+2])
		if offset < 0 {
			f.PC -= uint32(-offset)
		} else {
			f.PC += uint32(offset)
		}
		//f.PC = uint32(f.Text[f.PC+1])<<8 + uint32(f.Text[f.PC+2])
	} else {
		f.PC += 3
	}
	//f.Stack = f.Stack[:len(f.Stack)-2]
	//f.DataStack.Pop()
}
