package nbt

type ListTag []Tag

func NewListTag(val []Tag) ListTag {
	return ListTag(val)
}

func (ListTag) GetId() uint8 {
	return 9
}
