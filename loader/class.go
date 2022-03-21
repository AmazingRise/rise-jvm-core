package loader

type Class struct {
	This      string
	Methods   map[string]*Method
	Constants *ConstantPool
	Super     string
	Flags     uint16
}

func (c *Class) IsPublic() bool {
	return c.Flags&0x0001 > 0
}
