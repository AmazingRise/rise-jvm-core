package loader

import (
	"bytes"
	"io"
	"wasm-jvm/entity"
	"wasm-jvm/logger"
	"wasm-jvm/utils"
)

type ClassLoader struct {
	reader *utils.Reader
	//file   io.Reader
	class *entity.Class
}

// CreateLoader Initialization
func CreateLoader() *ClassLoader {
	return &ClassLoader{}
}

// Load To load the class
func (l *ClassLoader) Load(classFile io.Reader) *entity.Class {
	//l.file = classFile
	l.reader = utils.CreateReader(classFile)
	l.class = &entity.Class{}
	if !bytes.Equal(l.readBytes(4), []byte{0xCA, 0xFE, 0xBA, 0xBE}) { // magic number
		logger.Errorln("invalid java class file")
	}
	l.loadMeta()
	return l.class
}

// loadMeta To load meta information of the class file
/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
func (l *ClassLoader) loadMeta() {
	c := l.class
	l.readBytes(2)          // minor version
	major := l.readBytes(2) // major version
	logger.Infoln("Java Version: ", majorToInt(major))

	constantPoolCount := l.u2() - 1
	// in JVM specification, here should be minus 1
	logger.Infoln("Constant Pool Count: ", constantPoolCount)
	l.readConstantPool(constantPoolCount)
	p := c.Constants // constant pool

	c.Flags = l.u2() // access_flags

	thisIdx := l.u2()                     // this_class
	c.This = p.GetClassNameByIdx(thisIdx) // get the name of this

	superIdx := l.u2()                      // super_class
	c.Super = p.GetClassNameByIdx(superIdx) // get the name of super
	iCount := l.u2()                        // interfaces_count
	// interfaces
	logger.Infoln("Interfaces: ", iCount)
	for i := 0; i < int(iCount); i++ {
		l.u2()
	}

	fCount := l.u2() // fields_count
	// fields
	logger.Infoln("Fields: ", fCount)
	l.readFields(fCount)

	mCount := l.u2() // methods_count
	// methods
	logger.Infoln("Methods: ", mCount)
	l.readMethods(mCount)

	aCount := l.u2() // attributes_count
	// attributes
	l.readAttributes(aCount)
}

// Utils

// From https://zserge.com/posts/jvm/
func (l *ClassLoader) u1() uint8  { return l.readBytes(1)[0] }
func (l *ClassLoader) u2() uint16 { return l.reader.U2() }
func (l *ClassLoader) u4() uint32 { return l.reader.U4() }
func (l *ClassLoader) u8() uint64 { return l.reader.U8() }

func (l *ClassLoader) readBytes(n int) []byte {
	return l.reader.ReadBytes(n)
}

func majorToInt(bytes []byte) int {
	return int(bytes[1] - 44)
}
