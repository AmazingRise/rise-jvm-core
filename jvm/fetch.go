package jvm

import (
	"bytes"
	"rise-jvm-core/entity"
	"rise-jvm-core/utils"
)

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

func Load(raw []byte) *entity.ByteCode {
	code := &entity.ByteCode{}
	r := utils.CreateReader(bytes.NewReader(raw))
	code.MaxStack = r.U2()
	code.MaxLocals = r.U2()
	codeLen := r.U4()
	code.Text = r.ReadBytes(int(codeLen))
	exLen := r.U2()
	var i uint16
	for i = 0; i < exLen; i++ {
		code.ExceptionTable = append(code.ExceptionTable, entity.Exception{
			StartPc:   r.U2(),
			EndPc:     r.U2(),
			HandlerPc: r.U2(),
			CatchType: r.U2(),
		})
	}
	// TODO: Read attributes
	return code
}

/*
// Cut verify and cut the code into instructions
func Cut(code []byte) [][]byte {
	var seg [][]byte
	n := len(code)
	pos := 0
	for pos < n {
		op := code[pos]
		switch op {
		case OpILoad:
			seg = append(seg, []byte{code[pos], code[pos+1]})
			pos += 2
		case OpILoad0:
			fallthrough
		case OpILoad1:
			fallthrough
		case OpILoad2:
			fallthrough
		case OpILoad3:
			seg = append(seg, []byte{code[pos]})
			pos++
		case OpIAdd:
			seg = append(seg, []byte{code[pos]})
			pos++
		case OpIReturn:
			seg = append(seg, []byte{code[pos]})
			pos++
		default:
			// Should not reach here
			logger.Warnf("ins 0x%X not recognized", code[pos])
			logger.Errorln("should not reach here")
		}
	}
	return seg
}
*/
