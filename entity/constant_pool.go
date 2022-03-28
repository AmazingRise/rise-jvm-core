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

type ConstPool struct {
	Utf8Const  map[uint16]string
	ClassConst map[uint16]uint16

	MethodRefConst          map[uint16]Ref
	FieldRefConst           map[uint16]Ref
	InterfaceMethodRefConst map[uint16]Ref

	StrConst    map[uint16]uint16 // index
	IntConst    map[uint16]int
	FloatConst  map[uint16]float32
	LongConst   map[uint16]int64
	DoubleConst map[uint16]float64

	NameTypeConst map[uint16]NameType
}

type Ref struct {
	// TODO: Cache here
	ClassIdx    uint16
	NameTypeIdx uint16
}

type NameType struct {
	NameIdx uint16
	DescIdx uint16
}

// GetUtf8Const get an utf8 constant from constant pool by its index
func (p *ConstPool) GetUtf8Const(idx uint16) string {
	return p.Utf8Const[idx]
}

// GetClassName get a class name from constant pool by its index
func (p *ConstPool) GetClassName(idx uint16) string {
	nameIdx := p.ClassConst[idx]
	return p.GetUtf8Const(nameIdx)
}

// GetMethodRef get the class, name and description of a method, from its index in constant pool
func (p *ConstPool) GetMethodRef(idx uint16) (class string, name string, desc string) {
	methodRef := p.MethodRefConst[idx]
	class = p.GetClassName(methodRef.ClassIdx)
	name, desc = p.GetNameType(methodRef.NameTypeIdx)
	return
}

// GetFieldRef get the class, name and description of a field, from its index in constant pool
func (p *ConstPool) GetFieldRef(idx uint16) (class string, name string, desc string) {
	methodRef := p.FieldRefConst[idx]
	class = p.GetClassName(methodRef.ClassIdx)
	name, desc = p.GetNameType(methodRef.NameTypeIdx)
	return
}

// GetInterfaceMethodRef get the class, name and description of an interface method, from its index in constant pool
func (p *ConstPool) GetInterfaceMethodRef(idx uint16) (class string, name string, desc string) {
	methodRef := p.InterfaceMethodRefConst[idx]
	class = p.GetClassName(methodRef.ClassIdx)
	name, desc = p.GetNameType(methodRef.NameTypeIdx)
	return
}

func (p *ConstPool) GetNameType(idx uint16) (name string, desc string) {
	nameType := p.NameTypeConst[idx]
	name = p.GetUtf8Const(nameType.NameIdx)
	desc = p.GetUtf8Const(nameType.DescIdx)
	return
}

func (p *ConstPool) GetInt(idx uint16) int {
	return p.IntConst[idx]
}

func (p *ConstPool) GetStr(idx uint16) string {
	return p.GetUtf8Const(p.StrConst[idx])
}

func (p *ConstPool) GetLong(idx uint16) int64 {
	return p.LongConst[idx]
}

func (p *ConstPool) GetFloat(idx uint16) float32 {
	return p.FloatConst[idx]
}

func (p *ConstPool) GetDouble(idx uint16) float64 {
	return p.DoubleConst[idx]
}
