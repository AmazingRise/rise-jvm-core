package entity

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

func (p *ConstantPool) GetUtf8Constant(idx uint16) string {
	return p.Utf8Constants[idx]
}

func (p *ConstantPool) GetClassNameByIdx(idx uint16) string {
	nameIdx := p.ClassConstants[idx]
	return p.GetUtf8Constant(nameIdx)
}
