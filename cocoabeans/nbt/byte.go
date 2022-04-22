package nbt

type ByteTag int8

func NewByteTag(val int8) ByteTag {
	return ByteTag(val)
}

func (ByteTag) GetId() uint8 {
	return 1
}

func (tag ByteTag) AsByte() *ByteTag {
	return &tag
}

func (tag ByteTag) AsShort() *ShortTag {
	return nil
}

func (tag ByteTag) AsInt() *IntTag {
	return nil
}

func (tag ByteTag) AsLong() *LongTag {
	return nil
}

func (tag ByteTag) AsFloat() *FloatTag {
	return nil
}

func (tag ByteTag) AsDouble() *DoubleTag {
	return nil
}

func (tag ByteTag) AsByteArray() *ByteArrayTag {
	return nil
}

func (tag ByteTag) AsString() *StringTag {
	return nil
}

func (tag ByteTag) AsList() *ListTag {
	return nil
}

func (tag ByteTag) AsComound() *CompoundTag {
	return nil
}

func (tag ByteTag) AsIntArray() *IntArrayTag {
	return nil
}

func (tag ByteTag) AsLongArray() *LongArrayTag {
	return nil
}
