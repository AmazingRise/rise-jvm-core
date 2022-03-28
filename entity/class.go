package entity

type Class struct {
	This      string
	Methods   map[string]*Method
	Constants *ConstPool
	Super     string
	Flags     uint16
}

func (c *Class) IsPublic() bool {
	return c.Flags&0x0001 > 0
}
