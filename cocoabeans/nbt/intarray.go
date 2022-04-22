package nbt

type IntArrayTag []int32

func NewIntArrayTag(val []int32) IntArrayTag {
	return IntArrayTag(val)
}

func (IntArrayTag) GetId() uint8 {
	return 11
}
