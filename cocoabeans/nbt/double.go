package nbt

type DoubleTag float64

func NewDoubleTag(val float64) DoubleTag {
	return DoubleTag(val)
}

func (DoubleTag) GetId() uint8 {
	return 6
}
