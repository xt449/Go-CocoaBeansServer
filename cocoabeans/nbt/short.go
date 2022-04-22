package nbt

type ShortTag int16

func NewShortTag(val int16) ShortTag {
	return ShortTag(val)
}

func (ShortTag) GetId() uint8 {
	return 2
}
