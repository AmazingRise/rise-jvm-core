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
		case OpIAdd:
			result := t.Stack[0].(int) + t.Stack[1].(int)
			t.Stack = t.Stack[1:]
			t.Stack[0] = result
		case OpIReturn:
			return t.Stack[0]
		default:
			// Should not reach here
			logger.Warnf("ins 0x%X not recognized", opcode)
			logger.Errorln("should not reach here")
		}
		t.PC++
	}
	return nil
}
