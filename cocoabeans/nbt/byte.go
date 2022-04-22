package nbt

type ByteTag int8

func NewByteTag(val int8) ByteTag {
	return ByteTag(val)
}

func (ByteTag) GetId() uint8 {
	return 1
}
