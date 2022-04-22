package nbt

type CompoundTag map[string]Tag

func NewCompoundTag() CompoundTag {
	return CompoundTag(make(map[string]Tag))
}

func (CompoundTag) GetId() uint8 {
	return 10
}

func (tag *CompoundTag) AsByte() *ByteTag {
	return nil
}

func (tag *CompoundTag) AsShort() *ShortTag {
	return nil
}

func (tag *CompoundTag) AsInt() *IntTag {
	return nil
}

func (tag *CompoundTag) AsLong() *LongTag {
	return nil
}

func (tag *CompoundTag) AsFloat() *FloatTag {
	return nil
}

func (tag *CompoundTag) AsDouble() *DoubleTag {
	return nil
}

func (tag *CompoundTag) AsByteArray() *ByteArrayTag {
	return nil
}

func (tag *CompoundTag) AsString() *StringTag {
	return nil
}

func (tag *CompoundTag) AsList() *ListTag {
	return nil
}

func (tag *CompoundTag) AsComound() *CompoundTag {
	return tag
}

func (tag *CompoundTag) AsIntArray() *IntArrayTag {
	return nil
}

func (tag *CompoundTag) AsLongArray() *LongArrayTag {
	return nil
}
