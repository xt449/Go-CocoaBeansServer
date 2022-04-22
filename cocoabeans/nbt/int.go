package nbt

type IntTag int32

func NewIntTag(val int32) IntTag {
	return IntTag(val)
}

func (IntTag) GetId() uint8 {
	return 3
}
