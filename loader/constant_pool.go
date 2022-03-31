package loader

import (
	"rise-jvm-core/entity"
	"rise-jvm-core/logger"
)

func (l *Loader) readConstantPool(count uint16) {
	pool := entity.ConstPool{
		Utf8Const:               make(map[uint16]string),
		ClassConst:              make(map[uint16]uint16),
		MethodRefConst:          make(map[uint16]entity.Ref),
		FieldRefConst:           make(map[uint16]entity.Ref),
		InterfaceMethodRefConst: make(map[uint16]entity.Ref),
		StrConst:                make(map[uint16]uint16),
		IntConst:                make(map[uint16]int),
		FloatConst:              make(map[uint16]float32),
		LongConst:               make(map[uint16]int64),
		DoubleConst:             make(map[uint16]float64),
		NameTypeConst:           make(map[uint16]entity.NameType),
	}

	var i uint16
	for i = 1; i <= count; i++ {
		tag := l.u1()
		//logger.Infof("Constant #%d (tag: %d)", i, tag)
		switch tag {
		case entity.ConstantClass:
			// Class info
			//nameIdx = l.u2() // name_index
			pool.ClassConst[i] = l.u2()
		case entity.ConstantMethodref:
			ref := entity.Ref{
				ClassIdx:    l.u2(),
				NameTypeIdx: l.u2(),
			}
			pool.MethodRefConst[i] = ref
			//logger.Infof("Method ref class #%d, name and type #%d", ref.ClassIdx, ref.NameTypeIdx)
		case entity.ConstantInterfacemethodref:
			ref := entity.Ref{
				ClassIdx:    l.u2(),
				NameTypeIdx: l.u2(),
			}
			pool.InterfaceMethodRefConst[i] = ref
		case entity.ConstantFieldref:
			ref := entity.Ref{
				ClassIdx:    l.u2(),
				NameTypeIdx: l.u2(),
			}
			pool.FieldRefConst[i] = ref
		case entity.ConstantString:
			pool.StrConst[i] = l.u2()
		case entity.ConstantInteger:
			pool.IntConst[i] = int(l.u4())
		case entity.ConstantFloat:
			// pool.FloatConst[i] = float32(l.u4())
			l.u4()
		case entity.ConstantLong:
			pool.LongConst[i] = int64(l.u8())
		case entity.ConstantDouble:
			l.readBytes(4) // high bytes
			l.readBytes(4) // low bytes
		case entity.ConstantNameandtype:
			pool.NameTypeConst[i] = entity.NameType{
				NameIdx: l.u2(),
				DescIdx: l.u2(),
			}
			//logger.Infof("Name and types: name_index #%d, desc_index #%d", pool.NameTypeConst[i].NameIdx, pool.NameTypeConst[i].DescIdx)
		case entity.ConstantUtf8:
			length := l.u2() // length
			stringConst := l.readBytes(int(length))
			pool.Utf8Const[i] = string(stringConst)
			//logger.Infof("#%d string(%d): %s", i, length, string(stringConst))
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
