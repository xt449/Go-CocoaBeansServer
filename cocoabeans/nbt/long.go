package nbt

type LongTag int64

func NewLongTag(val int64) LongTag {
	return LongTag(val)
}

func (LongTag) GetId() uint8 {
	return 4
}

func (tag *LongTag) AsByte() *ByteTag {
	return nil
}

func (tag *LongTag) AsShort() *ShortTag {
	return nil
}

func (tag *LongTag) AsInt() *IntTag {
	return nil
}

func (tag *LongTag) AsLong() *LongTag {
	return tag
}

func (tag *LongTag) AsFloat() *FloatTag {
	return nil
}

func (tag *LongTag) AsDouble() *DoubleTag {
	return nil
}

func (tag *LongTag) AsByteArray() *ByteArrayTag {
	return nil
}

func (tag *LongTag) AsString() *StringTag {
	return nil
}

func (tag *LongTag) AsList() *ListTag {
	return nil
}

func (tag *LongTag) AsComound() *CompoundTag {
	return nil
}

func (tag *LongTag) AsIntArray() *IntArrayTag {
	return nil
}

func (tag *LongTag) AsLongArray() *LongArrayTag {
	return nil
}
