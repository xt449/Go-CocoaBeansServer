package nbt

type LongArrayTag []int64

func NewLongArrayTag(val []int64) LongArrayTag {
	return LongArrayTag(val)
}

func (LongArrayTag) GetId() uint8 {
	return 12
}

func (tag *LongArrayTag) AsByte() *ByteTag {
	return nil
}

func (tag *LongArrayTag) AsShort() *ShortTag {
	return nil
}

func (tag *LongArrayTag) AsInt() *IntTag {
	return nil
}

func (tag *LongArrayTag) AsLong() *LongTag {
	return nil
}

func (tag *LongArrayTag) AsFloat() *FloatTag {
	return nil
}

func (tag *LongArrayTag) AsDouble() *DoubleTag {
	return nil
}

func (tag *LongArrayTag) AsByteArray() *ByteArrayTag {
	return nil
}

func (tag *LongArrayTag) AsString() *StringTag {
	return nil
}

func (tag *LongArrayTag) AsList() *ListTag {
	return nil
}

func (tag *LongArrayTag) AsComound() *CompoundTag {
	return nil
}

func (tag *LongArrayTag) AsIntArray() *IntArrayTag {
	return nil
}

func (tag *LongArrayTag) AsLongArray() *LongArrayTag {
	return tag
}
