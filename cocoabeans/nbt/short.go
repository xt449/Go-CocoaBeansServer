package nbt

type ShortTag int16

func NewShortTag(val int16) ShortTag {
	return ShortTag(val)
}

func (ShortTag) GetId() uint8 {
	return 2
}

func (tag *ShortTag) AsByte() *ByteTag {
	return nil
}

func (tag *ShortTag) AsShort() *ShortTag {
	return tag
}

func (tag *ShortTag) AsInt() *IntTag {
	return nil
}

func (tag *ShortTag) AsLong() *LongTag {
	return nil
}

func (tag *ShortTag) AsFloat() *FloatTag {
	return nil
}

func (tag *ShortTag) AsDouble() *DoubleTag {
	return nil
}

func (tag *ShortTag) AsByteArray() *ByteArrayTag {
	return nil
}

func (tag *ShortTag) AsString() *StringTag {
	return nil
}

func (tag *ShortTag) AsList() *ListTag {
	return nil
}

func (tag *ShortTag) AsComound() *CompoundTag {
	return nil
}

func (tag *ShortTag) AsIntArray() *IntArrayTag {
	return nil
}

func (tag *ShortTag) AsLongArray() *LongArrayTag {
	return nil
}
