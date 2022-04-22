package nbt

type FloatTag float32

func NewFloatTag(val float32) FloatTag {
	return FloatTag(val)
}

func (FloatTag) GetId() uint8 {
	return 5
}

func (tag *FloatTag) AsByte() *ByteTag {
	return nil
}

func (tag *FloatTag) AsShort() *ShortTag {
	return nil
}

func (tag *FloatTag) AsInt() *IntTag {
	return nil
}

func (tag *FloatTag) AsLong() *LongTag {
	return nil
}

func (tag *FloatTag) AsFloat() *FloatTag {
	return tag
}

func (tag *FloatTag) AsDouble() *DoubleTag {
	return nil
}

func (tag *FloatTag) AsByteArray() *ByteArrayTag {
	return nil
}

func (tag *FloatTag) AsString() *StringTag {
	return nil
}

func (tag *FloatTag) AsList() *ListTag {
	return nil
}

func (tag *FloatTag) AsComound() *CompoundTag {
	return nil
}

func (tag *FloatTag) AsIntArray() *IntArrayTag {
	return nil
}

func (tag *FloatTag) AsLongArray() *LongArrayTag {
	return nil
}
