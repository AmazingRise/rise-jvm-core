package entity

type Object struct {
	Proto *Class
	Data  map[string]interface{}
}

func CreateObject(Proto *Class) *Object {
	return &Object{
		Proto: Proto,
		Data:  make(map[string]interface{}),
	}
}

func (o *Object) Get(name string) interface{} {
	// TODO: Exception process
	return o.Data[name]
}

func (o *Object) Set(name string, value interface{}) {
	// TODO: Safety check
	o.Data[name] = value
}
