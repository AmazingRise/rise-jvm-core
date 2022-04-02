package entity

const (
	ClassPublic     = 0x1
	ClassPrivate    = 0x2
	ClassProtected  = 0x4
	ClassStatic     = 0x8
	ClassFinal      = 0x10
	ClassInterface  = 0x200
	ClassAbstract   = 0x400
	ClassSynthetic  = 0x1000
	ClassAnnotation = 0x2000
	ClassEnum       = 0x4000
)

type Class struct {
	Name      string
	Methods   map[string][]*Method
	Constants *ConstPool
	Super     string
	Flags     uint16
	fields    map[string]interface{}
}

func (c *Class) IsPublic() bool {
	return c.Flags&ClassPublic > 0
}

func (c *Class) SetStaticField(name string, value interface{}) {
	// TODO: Exception process
	c.fields[name] = value
}

func (c *Class) GetStaticField(name string) interface{} {
	if c == nil {
		// Fake method
		return nil
	}
	return c.fields[name]
}
