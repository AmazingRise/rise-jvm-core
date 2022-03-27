package loader

import (
	"wasm-jvm/entity"
	"wasm-jvm/logger"
)

func (l *ClassLoader) readConstantPool(count uint16) {
	pool := entity.ConstantPool{}
	pool.Utf8Constants = make(map[uint16]string)
	pool.ClassConstants = make(map[uint16]uint16)
	var i uint16
	for i = 1; i <= count; i++ {
		tag := l.u1()
		logger.Infof("Constant #%d (tag: %d)", i, tag)
		switch tag {
		case entity.ConstantClass:
			// Class info
			//nameIdx = l.u2() // name_index
			pool.ClassConstants[i] = l.u2()
		case entity.ConstantMethodref:
			fallthrough
		case entity.ConstantInterfacemethodref:
			fallthrough
		case entity.ConstantFieldref:
			l.readBytes(2) // class_index
			l.readBytes(2) // name_and_type_index
		case entity.ConstantString:
			l.readBytes(2) // string_index
		case entity.ConstantInteger:
			l.readBytes(4) // bytes
		case entity.ConstantFloat:
			l.readBytes(4) // bytes
		case entity.ConstantLong:
			fallthrough
		case entity.ConstantDouble:
			l.readBytes(4) // high bytes
			l.readBytes(4) // low bytes
		case entity.ConstantNameandtype:
			nameIdx := l.u2() // name_index
			descIdx := l.u2() // descriptor_index
			logger.Infof("Name and types: name_index #%d, desc_index #%d", nameIdx, descIdx)
		case entity.ConstantUtf8:
			length := l.u2() // length
			stringConst := l.readBytes(int(length))
			pool.Utf8Constants[i] = string(stringConst)
			logger.Infof("#%d string(%d): %s", i, length, string(stringConst))
		case entity.CONSTANT_InvokeDynamic:
			l.u2()
			l.u2()
		case entity.CONSTANT_MethodHandle:
			l.u1()
			l.u2()
		default:
			logger.Errorln("Should not reach here.")
		}
	}
	// after reading constant pool:
	//fmt.Infoln(nameIdx)
	l.class.Constants = &pool
}
