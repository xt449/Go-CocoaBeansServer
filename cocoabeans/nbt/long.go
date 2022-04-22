package nbt

type LongTag int64

func NewLongTag(val int64) LongTag {
	return LongTag(val)
}

func (LongTag) GetId() uint8 {
	return 4
}
