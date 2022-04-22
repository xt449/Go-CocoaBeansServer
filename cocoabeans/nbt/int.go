package nbt

type IntTag int32

func NewIntTag(val int32) IntTag {
	return IntTag(val)
}

func (IntTag) GetId() uint8 {
	return 3
}

func (tag *IntTag) AsByte() *ByteTag {
	return nil
}

func (tag *IntTag) AsShort() *ShortTag {
	return nil
}

func (tag *IntTag) AsInt() *IntTag {
	return tag
}

func (tag *IntTag) AsLong() *LongTag {
	return nil
}

func (tag *IntTag) AsFloat() *FloatTag {
	return nil
}

func (tag *IntTag) AsDouble() *DoubleTag {
	return nil
}

func (tag *IntTag) AsByteArray() *ByteArrayTag {
	return nil
}

func (tag *IntTag) AsString() *StringTag {
	return nil
}

func (tag *IntTag) AsList() *ListTag {
	return nil
}

func (tag *IntTag) AsComound() *CompoundTag {
	return nil
}

func (tag *IntTag) AsIntArray() *IntArrayTag {
	return nil
}

func (tag *IntTag) AsLongArray() *LongArrayTag {
	return nil
}
