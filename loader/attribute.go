package loader

import (
	"wasm-jvm/entity"
	"wasm-jvm/logger"
)

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

func (l *ClassLoader) readAttributes(count uint16) []entity.Attribute {
	var result []entity.Attribute
	var i uint16
	for i = 0; i < count; i++ {
		nameIdx := l.u2()
		name := l.class.Constants.GetUtf8Constant(nameIdx)
		logger.Infoln("Attribute:", name)
		aLen := l.u4() // attribute length
		bytes := l.readBytes(int(aLen))
		result = append(result, entity.Attribute{name, bytes})
	}
	return result
}
