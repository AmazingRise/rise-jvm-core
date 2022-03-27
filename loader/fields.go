package loader

/*
field_info {
	u2             access_flags;
	u2             name_index;
	u2             descriptor_index;
	u2             attributes_count;
	attribute_info attributes[attributes_count];
}
*/

func (l *Loader) readFields(count uint16) {
	var i uint16
	for i = 0; i < count; i++ {
		l.u2()                   // access flags
		l.u2()                   // name index
		l.u2()                   // descriptor index
		aCount := l.u2()         // attributes_count
		l.ReadAttributes(aCount) // read attributes
	}

}
