package loader

import (
	"rise-jvm-core/entity"
	"rise-jvm-core/logger"
)

func (l *Loader) readConstantPool(count uint16) {
	pool := entity.ConstPool{
		Consts: make([]interface{}, count+1),
	}

	var i uint16
	var strIdx []uint16
	for i = 1; i <= count; i++ {
		tag := l.u1()
		//logger.Infof("Constant #%d (tag: %d)", i, tag)
		switch tag {
		case entity.ConstantClass:
			// Class info
			//nameIdx = l.u2() // name_index
			pool.Consts[i] = l.u2()
		case entity.ConstantMethodref:
			ref := entity.Ref{
				ClassIdx:    l.u2(),
				NameTypeIdx: l.u2(),
			}
			pool.Consts[i] = ref
			//logger.Infof("Method ref class #%d, name and type #%d", ref.ClassIdx, ref.NameTypeIdx)
		case entity.ConstantInterfacemethodref:
			ref := entity.Ref{
				ClassIdx:    l.u2(),
				NameTypeIdx: l.u2(),
			}
			pool.Consts[i] = ref
		case entity.ConstantFieldref:
			ref := entity.Ref{
				ClassIdx:    l.u2(),
				NameTypeIdx: l.u2(),
			}
			pool.Consts[i] = ref
		case entity.ConstantString:
			idx := l.u2()
			strIdx = append(strIdx, i)
			pool.Consts[i] = idx
		case entity.ConstantInteger:
			pool.Consts[i] = int(l.u4())
		case entity.ConstantFloat:
			// pool.FloatConst[i] = float32(l.u4())
			pool.Consts[i] = float32(l.u4())
		case entity.ConstantLong:
			pool.Consts[i] = int64(l.u8())
		case entity.ConstantDouble:
			l.readBytes(4) // high bytes
			l.readBytes(4) // low bytes
		case entity.ConstantNameandtype:
			pool.Consts[i] = entity.NameType{
				NameIdx: l.u2(),
				DescIdx: l.u2(),
			}
			//logger.Infof("Name and types: name_index #%d, desc_index #%d", pool.NameTypeConst[i].NameIdx, pool.NameTypeConst[i].DescIdx)
		case entity.ConstantUtf8:
			length := l.u2() // length
			stringConst := l.readBytes(int(length))
			pool.Consts[i] = string(stringConst)
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
	for _, idx := range strIdx {
		pool.Consts[idx] = pool.GetUtf8Const(pool.Consts[idx].(uint16))
	}
	l.class.Constants = &pool
}
