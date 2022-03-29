package jvm

import (
	"wasm-jvm/entity"
	"wasm-jvm/logger"
)

type Frame struct {
	entity.ByteCode
	PC     uint32
	Stack  []interface{}
	Locals []interface{}
	//Ins    [][]byte

	MethodName string
	State      int

	This *entity.Class
}

const (
	FrameReady = 0
	FrameExit  = 1
	FramePush  = 2
)

// Exec execute the frame
func (f *Frame) Exec() []interface{} {
	n := len(f.Text)
	if int(f.PC) < n {
		opcode := f.Text[f.PC]
		//fmt.Printf("0x%X", opcode)
		switch opcode {
		case OpILoad:
			idx := f.Text[f.PC+1]
			f.Stack = append(f.Stack, f.Locals[idx])
			f.PC++
		case OpILoad0:
			f.Stack = append(f.Stack, f.Locals[0])
		case OpILoad1:
			f.Stack = append(f.Stack, f.Locals[1])
		case OpILoad2:
			f.Stack = append(f.Stack, f.Locals[2])
		case OpILoad3:
			f.Stack = append(f.Stack, f.Locals[3])
		case OpIStore:
			idx := f.Text[f.PC+1]
			f.Locals[idx] = f.Stack[0]
			f.Stack = f.Stack[1:]
			f.PC++
		case OpIStore0:
			f.Locals[0] = f.Stack[0]
			f.Stack = f.Stack[1:]
		case OpIStore1:
			f.Locals[1] = f.Stack[0]
			f.Stack = f.Stack[1:]
		case OpIStore2:
			f.Locals[2] = f.Stack[0]
			f.Stack = f.Stack[1:]
		case OpIStore3:
			f.Locals[3] = f.Stack[0]
			f.Stack = f.Stack[1:]
		case OpIAdd:
			result := f.Stack[0].(int) + f.Stack[1].(int)
			f.Stack = f.Stack[1:]
			f.Stack[0] = result
		case OpIMul:
			result := f.Stack[0].(int) * f.Stack[1].(int)
			f.Stack = f.Stack[1:]
			f.Stack[0] = result
		case OpIReturn:
			f.State = FrameExit
			f.PC++
			return f.Stack
		case OpIConst1:
			fallthrough
		case OpIConst2:
			fallthrough
		case OpIConst3:
			fallthrough
		case OpIConst4:
			fallthrough
		case OpIConst5:
			fallthrough
		case OpIConst0:
			fallthrough
		case OpIConstM1:
			f.Stack = append(f.Stack, opcode-3)
		case OpBiPush:
			// Set it to int
			f.Stack = append(f.Stack, int(f.Text[f.PC+1]))
			f.PC++
		case OpSiPush:
			// (byte1 << 8) | byte2
			short := int(f.Text[f.PC+1])<<8 + int(f.Text[f.PC+2])
			f.Stack = append(f.Stack, short)
			f.PC += 2
		case OpInvokeStatic:
			// Invoke a static method
			idx := uint16(f.Text[f.PC+1])<<8 + uint16(f.Text[f.PC+2])
			f.State = FramePush
			f.PC += 3
			return []interface{}{idx}
		case OpReturn:
			f.State = FrameExit
		default:
			// Should not reach here
			logger.Warnf("ins 0x%X not recognized", opcode)
			logger.Errorln("should not reach here")
		}
		f.PC++
	}
	return nil
}
