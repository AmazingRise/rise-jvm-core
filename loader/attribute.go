package loader

import (
	"rise-jvm-core/entity"
	"rise-jvm-core/logger"
)

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

func (l *Loader) ReadAttributes(count uint16) []entity.Attribute {
	var result []entity.Attribute
	var i uint16
	for i = 0; i < count; i++ {
		nameIdx := l.u2()
		name := l.class.Constants.GetUtf8Const(nameIdx)
		logger.Infoln("Attribute:", name)
		aLen := l.u4() // attribute length
		bytes := l.readBytes(int(aLen))
		result = append(result, entity.Attribute{Name: name, Bytes: bytes})
	}
	return result
}
