package jvm

import (
	"wasm-jvm/entity"
	"wasm-jvm/logger"
)

const (
	ThreadReady   = 0
	ThreadBlocked = 1
	ThreadDead    = 2
)

type Thread struct {
	entity.ByteCode
	PC     uint32
	Stack  []interface{}
	Locals []interface{}
	//Ins    [][]byte
	State int
}

func (t *Thread) Exec() interface{} {
	n := len(t.Text)
	for int(t.PC) < n {
		opcode := t.Text[t.PC]
		switch opcode {
		case OpILoad:
			idx := t.Text[t.PC+1]
			t.Stack = append(t.Stack, t.Locals[idx])
			t.PC++
		case OpILoad0:
			t.Stack = append(t.Stack, t.Locals[0])
		case OpILoad1:
			t.Stack = append(t.Stack, t.Locals[1])
		case OpILoad2:
			t.Stack = append(t.Stack, t.Locals[2])
		case OpILoad3:
			t.Stack = append(t.Stack, t.Locals[3])
		case OpIStore:
			idx := t.Text[t.PC+1]
			t.Locals[idx] = t.Stack[0]
			t.Stack = t.Stack[1:]
			t.PC++
		case OpIStore0:
			t.Locals[0] = t.Stack[0]
			t.Stack = t.Stack[1:]
		case OpIStore1:
			t.Locals[1] = t.Stack[0]
			t.Stack = t.Stack[1:]
		case OpIStore2:
			t.Locals[2] = t.Stack[0]
			t.Stack = t.Stack[1:]
		case OpIStore3:
			t.Locals[3] = t.Stack[0]
			t.Stack = t.Stack[1:]
		case OpIAdd:
			result := t.Stack[0].(int) + t.Stack[1].(int)
			t.Stack = t.Stack[1:]
			t.Stack[0] = result
		case OpIMul:
			result := t.Stack[0].(int) * t.Stack[1].(int)
			t.Stack = t.Stack[1:]
			t.Stack[0] = result
		case OpIReturn:
			return t.Stack[0]
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
			t.Stack = append(t.Stack, opcode-3)
		case OpBiPush:
			// Set it to int
			t.Stack = append(t.Stack, int(t.Text[t.PC+1]))
			t.PC++
		case OpSiPush:
			// (byte1 << 8) | byte2
			short := int(t.Text[t.PC+1])<<8 + int(t.Text[t.PC+2])
			t.Stack = append(t.Stack, short)
			t.PC += 2
		default:
			// Should not reach here
			logger.Warnf("ins 0x%X not recognized", opcode)
			logger.Errorln("should not reach here")
		}
		t.PC++
	}
	return nil
}
