package nbt

type StringTag string

func NewStringTag(val string) StringTag {
	return StringTag(val)
}

func (StringTag) GetId() uint8 {
	return 8
}

func (tag *StringTag) AsByte() *ByteTag {
	return nil
}

func (tag *StringTag) AsShort() *ShortTag {
	return nil
}

func (tag *StringTag) AsInt() *IntTag {
	return nil
}

func (tag *StringTag) AsLong() *LongTag {
	return nil
}

func (tag *StringTag) AsFloat() *FloatTag {
	return nil
}

func (tag *StringTag) AsDouble() *DoubleTag {
	return nil
}

func (tag *StringTag) AsByteArray() *ByteArrayTag {
	return nil
}

func (tag *StringTag) AsString() *StringTag {
	return tag
}

func (tag *StringTag) AsList() *ListTag {
	return nil
}

func (tag *StringTag) AsComound() *CompoundTag {
	return nil
}

func (tag *StringTag) AsIntArray() *IntArrayTag {
	return nil
}

func (tag *StringTag) AsLongArray() *LongArrayTag {
	return nil
}
