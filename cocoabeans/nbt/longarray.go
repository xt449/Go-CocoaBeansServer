package nbt

type LongArrayTag []int64

func NewLongArrayTag(val []int64) LongArrayTag {
	return LongArrayTag(val)
}

func (LongArrayTag) GetId() uint8 {
	return 12
}
