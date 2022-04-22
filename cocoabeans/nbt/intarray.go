package nbt

type IntArrayTag []int32

func NewIntArrayTag(val []int32) IntArrayTag {
	return IntArrayTag(val)
}

func (IntArrayTag) GetId() uint8 {
	return 11
}

func (tag *IntArrayTag) AsByte() *ByteTag {
	return nil
}

func (tag *IntArrayTag) AsShort() *ShortTag {
	return nil
}

func (tag *IntArrayTag) AsInt() *IntTag {
	return nil
}

func (tag *IntArrayTag) AsLong() *LongTag {
	return nil
}

func (tag *IntArrayTag) AsFloat() *FloatTag {
	return nil
}

func (tag *IntArrayTag) AsDouble() *DoubleTag {
	return nil
}

func (tag *IntArrayTag) AsByteArray() *ByteArrayTag {
	return nil
}

func (tag *IntArrayTag) AsString() *StringTag {
	return nil
}

func (tag *IntArrayTag) AsList() *ListTag {
	return nil
}

func (tag *IntArrayTag) AsComound() *CompoundTag {
	return nil
}

func (tag *IntArrayTag) AsIntArray() *IntArrayTag {
	return tag
}

func (tag *IntArrayTag) AsLongArray() *LongArrayTag {
	return nil
}
