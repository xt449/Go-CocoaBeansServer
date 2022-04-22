package nbt

type DoubleTag float64

func NewDoubleTag(val float64) DoubleTag {
	return DoubleTag(val)
}

func (DoubleTag) GetId() uint8 {
	return 6
}

func (tag *DoubleTag) AsByte() *ByteTag {
	return nil
}

func (tag *DoubleTag) AsShort() *ShortTag {
	return nil
}

func (tag *DoubleTag) AsInt() *IntTag {
	return nil
}

func (tag *DoubleTag) AsLong() *LongTag {
	return nil
}

func (tag *DoubleTag) AsFloat() *FloatTag {
	return nil
}

func (tag *DoubleTag) AsDouble() *DoubleTag {
	return tag
}

func (tag *DoubleTag) AsByteArray() *ByteArrayTag {
	return nil
}

func (tag *DoubleTag) AsString() *StringTag {
	return nil
}

func (tag *DoubleTag) AsList() *ListTag {
	return nil
}

func (tag *DoubleTag) AsComound() *CompoundTag {
	return nil
}

func (tag *DoubleTag) AsIntArray() *IntArrayTag {
	return nil
}

func (tag *DoubleTag) AsLongArray() *LongArrayTag {
	return nil
}
