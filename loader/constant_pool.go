package loader

// Constant Type

const (
	ConstantClass              byte = 7
	ConstantFieldref           byte = 9
	ConstantMethodref          byte = 10
	ConstantInterfacemethodref byte = 11
	ConstantString             byte = 8
	ConstantInteger            byte = 3
	ConstantFloat              byte = 4
	ConstantLong               byte = 5
	ConstantDouble             byte = 6
	ConstantNameandtype        byte = 12
	ConstantUtf8               byte = 1
	CONSTANT_MethodHandle      byte = 15
	CONSTANT_MethodType        byte = 16
	CONSTANT_InvokeDynamic     byte = 18
)

type ConstantPool struct {
	Utf8Constants  map[uint16]string
	ClassConstants map[uint16]uint16
}

func (l *ClassLoader) readConstantPool(count uint16) {
	pool := ConstantPool{}
	pool.Utf8Constants = make(map[uint16]string)
	pool.ClassConstants = make(map[uint16]uint16)
	var i uint16
	for i = 1; i <= count; i++ {
		tag := l.u1()
		debugLog.Printf("Constant #%d (tag: %d)", i, tag)
		switch tag {
		case ConstantClass:
			// Class info
			//nameIdx = l.u2() // name_index
			pool.ClassConstants[i] = l.u2()
		case ConstantMethodref:
			fallthrough
		case ConstantInterfacemethodref:
			fallthrough
		case ConstantFieldref:
			l.readBytes(2) // class_index
			l.readBytes(2) // name_and_type_index
		case ConstantString:
			l.readBytes(2) // string_index
		case ConstantInteger:
			l.readBytes(4) // bytes
		case ConstantFloat:
			l.readBytes(4) // bytes
		case ConstantLong:
			fallthrough
		case ConstantDouble:
			l.readBytes(4) // high bytes
			l.readBytes(4) // low bytes
		case ConstantNameandtype:
			nameIdx := l.u2() // name_index
			descIdx := l.u2() // descriptor_index
			debugLog.Printf("Name and types: name_index #%d, desc_index #%d", nameIdx, descIdx)
		case ConstantUtf8:
			length := l.u2() // length
			stringConst := l.readBytes(int(length))
			pool.Utf8Constants[i] = string(stringConst)
			debugLog.Printf("#%d string(%d): %s", i, length, string(stringConst))
		case CONSTANT_InvokeDynamic:
			l.u2()
			l.u2()
		case CONSTANT_MethodHandle:
			l.u1()
			l.u2()
		default:
			errLog.Fatalf("Should not reach here.")
		}
	}
	// after reading constant pool:
	//fmt.Println(nameIdx)
	l.class.Constants = &pool
}

func (p *ConstantPool) getUtf8Constant(idx uint16) string {
	return p.Utf8Constants[idx]
}

func (p *ConstantPool) getClassNameByIdx(idx uint16) string {
	nameIdx := p.ClassConstants[idx]
	return p.getUtf8Constant(nameIdx)
}
