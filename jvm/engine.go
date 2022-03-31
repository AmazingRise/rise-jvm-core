package jvm

import (
	"rise-jvm-core/entity"
	"rise-jvm-core/logger"
)

// Exec execute the frame
func (v *VM) Exec(f *Frame) []interface{} {
	n := len(f.Text)
	if int(f.PC) < n {
		opcode := f.Text[f.PC]
		logger.Infoln("Stack:", f.Stack)
		logger.Infof("PC: %d, OpCode: %x", f.PC, opcode)

		switch opcode {
		case OpNop:
			break
		case OpLDC:
			idx := f.Text[f.PC+1]
			f.Stack = append(f.Stack, f.This.Constants.GetConst(uint16(idx)))
			f.PC++
		case OpALoad:
			fallthrough
		case OpILoad:
			idx := f.Text[f.PC+1]
			f.Stack = append(f.Stack, f.Locals[idx])
			f.PC++
		case OpALoad0:
			fallthrough
		case OpILoad0:
			f.Stack = append(f.Stack, f.Locals[0])
		case OpALoad1:
			fallthrough
		case OpILoad1:
			f.Stack = append(f.Stack, f.Locals[1])
		case OpALoad2:
			fallthrough
		case OpILoad2:
			f.Stack = append(f.Stack, f.Locals[2])
		case OpALoad3:
			fallthrough
		case OpILoad3:
			f.Stack = append(f.Stack, f.Locals[3])
		case OpAStore:
			fallthrough
		case OpIStore:
			idx := f.Text[f.PC+1]
			f.Locals[idx] = f.Stack[0]
			f.Stack = f.Stack[1:]
			f.PC++
		case OpAStore0:
			fallthrough
		case OpIStore0:
			f.Locals[0] = f.Stack[0]
			f.Stack = f.Stack[1:]
		case OpAStore1:
			fallthrough
		case OpIStore1:
			f.Locals[1] = f.Stack[0]
			f.Stack = f.Stack[1:]
		case OpAStore2:
			fallthrough
		case OpIStore2:
			f.Locals[2] = f.Stack[0]
			f.Stack = f.Stack[1:]
		case OpAStore3:
			fallthrough
		case OpIStore3:
			f.Locals[3] = f.Stack[0]
			f.Stack = f.Stack[1:]
		case OpDup:
			f.Stack = append(f.Stack, f.Stack[len(f.Stack)-1])
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
			f.Stack = append(f.Stack, int(opcode-3))
		case OpBiPush:
			// Set it to int
			f.Stack = append(f.Stack, int(f.Text[f.PC+1]))
			f.PC++
		case OpSiPush:
			// (byte1 << 8) | byte2
			short := int(f.Text[f.PC+1])<<8 + int(f.Text[f.PC+2])
			f.Stack = append(f.Stack, short)
			f.PC += 2
		case OpGetStatic:
			idx := uint16(f.Text[f.PC+1])<<8 + uint16(f.Text[f.PC+2])
			class, name, _ := f.This.Constants.GetFieldRef(idx)
			classRef := v.LocateClass(class)
			value := classRef.GetStaticField(name)
			f.Stack = append(f.Stack, value)
			f.PC += 2
		case OpGetField:
			idx := uint16(f.Text[f.PC+1])<<8 + uint16(f.Text[f.PC+2])
			// It should only be current object's field, so class is ignored
			_, name, _ := f.This.Constants.GetFieldRef(idx)
			// Pop the stack and get the object
			obj := f.Stack[len(f.Stack)-1].(*entity.Object)
			f.Stack[len(f.Stack)-1] = obj.Get(name)
			f.PC += 2
		case OpPutField:
			idx := uint16(f.Text[f.PC+1])<<8 + uint16(f.Text[f.PC+2])
			// It should only be current object's field, so class is ignored
			_, name, _ := f.This.Constants.GetFieldRef(idx)
			// Set the field
			obj := f.Stack[len(f.Stack)-2].(*entity.Object)
			value := f.Stack[len(f.Stack)-1]
			obj.Set(name, value)
			// Pop the objRef and value
			f.Stack = f.Stack[:len(f.Stack)-2]
			f.PC += 2
		case OpInvokeStatic:
			// Invoke a static method
			idx := uint16(f.Text[f.PC+1])<<8 + uint16(f.Text[f.PC+2])
			f.State = FramePush
			f.PC += 3 // directly return, so add 3
			return []interface{}{idx}
		case OpInvokeDynamic:
			// Invoke a dynamic method
			idx := uint16(f.Text[f.PC+1])<<8 + uint16(f.Text[f.PC+2])
			f.State = FramePush
			f.PC += 3 // directly return, so add 3
			return []interface{}{idx}
		case OpInvokeVirtual:
			// Invoke instance method
			idx := uint16(f.Text[f.PC+1])<<8 + uint16(f.Text[f.PC+2])
			f.State = FramePush
			f.PC += 3 // directly return, so add 3
			return []interface{}{idx}
		case OpInvokeSpecial:
			idx := uint16(f.Text[f.PC+1])<<8 + uint16(f.Text[f.PC+2])
			f.State = FramePush
			f.PC += 3 // directly return, so add 3
			return []interface{}{idx}
		case OpNew:
			idx := uint16(f.Text[f.PC+1])<<8 + uint16(f.Text[f.PC+2])
			className := f.This.Constants.GetClassName(idx)
			class := v.LocateClass(className)
			obj := entity.CreateObject(class)
			f.Stack = append(f.Stack, obj)
			f.PC += 2
		case OpIfICmpEq:
			f.condJmp(f.Stack[len(f.Stack)-2].(int) == f.Stack[len(f.Stack)-1].(int))
			return nil
		case OpIfICmpNe:
			f.condJmp(f.Stack[len(f.Stack)-2].(int) != f.Stack[len(f.Stack)-1].(int))
			return nil
		case OpIfICmpLt:
			f.condJmp(f.Stack[len(f.Stack)-2].(int) < f.Stack[len(f.Stack)-1].(int))
			return nil
		case OpIfICmpGe:
			f.condJmp(f.Stack[len(f.Stack)-2].(int) >= f.Stack[len(f.Stack)-1].(int))
			return nil
		case OpIfICmpGt:
			f.condJmp(f.Stack[len(f.Stack)-2].(int) > f.Stack[len(f.Stack)-1].(int))
			return nil
		case OpIfICmpLe:
			f.condJmp(f.Stack[len(f.Stack)-2].(int) <= f.Stack[len(f.Stack)-1].(int))
			return nil
		case OpGoto:
			offset := int16(f.Text[f.PC+1])<<8 + int16(f.Text[f.PC+2])
			if offset < 0 {
				f.PC -= uint32(-offset)
			} else {
				f.PC += uint32(offset)
			}
			return nil
		case OpIInc:
			idx := int(f.Text[f.PC+1])
			inc := int(f.Text[f.PC+2])
			f.Locals[idx] = f.Locals[idx].(int) + inc
			f.PC += 2
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
