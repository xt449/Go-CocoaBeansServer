package nbt

type ByteArrayTag []int8

func NewByteArrayTag(val []int8) ByteArrayTag {
	return ByteArrayTag(val)
}

func (ByteArrayTag) GetId() uint8 {
	return 7
}
