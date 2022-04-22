package nbt

type ListTag []Tag

func NewListTag(val []Tag) ListTag {
	return ListTag(val)
}

func (ListTag) GetId() uint8 {
	return 9
}

func (tag *ListTag) AsByte() *ByteTag {
	return nil
}

func (tag *ListTag) AsShort() *ShortTag {
	return nil
}

func (tag *ListTag) AsInt() *IntTag {
	return nil
}

func (tag *ListTag) AsLong() *LongTag {
	return nil
}

func (tag *ListTag) AsFloat() *FloatTag {
	return nil
}

func (tag *ListTag) AsDouble() *DoubleTag {
	return nil
}

func (tag *ListTag) AsByteArray() *ByteArrayTag {
	return nil
}

func (tag *ListTag) AsString() *StringTag {
	return nil
}

func (tag *ListTag) AsList() *ListTag {
	return tag
}

func (tag *ListTag) AsComound() *CompoundTag {
	return nil
}

func (tag *ListTag) AsIntArray() *IntArrayTag {
	return nil
}

func (tag *ListTag) AsLongArray() *LongArrayTag {
	return nil
}
