package nbt

type ByteArrayTag []int8

func NewByteArrayTag(val []int8) ByteArrayTag {
	return ByteArrayTag(val)
}

func (ByteArrayTag) GetId() uint8 {
	return 7
}

func (tag *ByteArrayTag) AsByte() *ByteTag {
	return nil
}

func (tag *ByteArrayTag) AsShort() *ShortTag {
	return nil
}

func (tag *ByteArrayTag) AsInt() *IntTag {
	return nil
}

func (tag *ByteArrayTag) AsLong() *LongTag {
	return nil
}

func (tag *ByteArrayTag) AsFloat() *FloatTag {
	return nil
}

func (tag *ByteArrayTag) AsDouble() *DoubleTag {
	return nil
}

func (tag *ByteArrayTag) AsByteArray() *ByteArrayTag {
	return tag
}

func (tag *ByteArrayTag) AsString() *StringTag {
	return nil
}

func (tag *ByteArrayTag) AsList() *ListTag {
	return nil
}

func (tag *ByteArrayTag) AsComound() *CompoundTag {
	return nil
}

func (tag *ByteArrayTag) AsIntArray() *IntArrayTag {
	return nil
}

func (tag *ByteArrayTag) AsLongArray() *LongArrayTag {
	return nil
}
