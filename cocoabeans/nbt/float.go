package nbt

type FloatTag float32

func NewFloatTag(val float32) FloatTag {
	return FloatTag(val)
}

func (FloatTag) GetId() uint8 {
	return 5
}
