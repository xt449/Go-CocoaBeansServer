package nbt

type StringTag string

func NewStringTag(val string) StringTag {
	return StringTag(val)
}

func (StringTag) GetId() uint8 {
	return 8
}
