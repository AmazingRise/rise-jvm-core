package entity

type Method struct {
	Name  string
	Flags uint16
	Desc  string
	Attrs []Attribute
	This  *Class
}

func (m *Method) IsPublic() bool {
	return m.Flags&0x0001 > 0
}

func (m *Method) IsStatic() bool {
	return m.Flags&0x0008 > 0
}
