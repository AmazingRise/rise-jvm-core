package entity

type Method struct {
	Name       string
	Flags      uint16
	Descriptor string
	Attrs      []Attribute
}

func (m *Method) IsPublic() bool {
	return m.Flags&0x0001 > 0
}

func (m *Method) IsStatic() bool {
	return m.Flags&0x0008 > 0
}
