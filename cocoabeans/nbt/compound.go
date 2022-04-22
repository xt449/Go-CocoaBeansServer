package nbt

type CompoundTag map[string]Tag

func NewCompoundTag() CompoundTag {
	return CompoundTag(make(map[string]Tag))
}

func (CompoundTag) GetId() uint8 {
	return 10
}
