package nbt

type Container interface {
	GetRoot() CompoundTag
}

type Tag interface {
	AsByte() *ByteTag
	AsShort() *ShortTag
	AsInt() *IntTag
	AsLong() *LongTag
	AsFloat() *FloatTag
	AsDouble() *DoubleTag
	AsByteArray() *ByteArrayTag
	AsString() *StringTag
	AsList() *ListTag
	AsComound() *CompoundTag
	AsIntArray() *IntArrayTag
	AsLongArray() *LongArrayTag
}

type TagId interface {
	GetId() uint8
}

type nbtContainer struct {
	root CompoundTag
}

func (container nbtContainer) GetRoot() CompoundTag {
	return container.root
}

func (container nbtContainer) Test() Tag {
	return container.root.AsLongArray()
}

func NewContainer() Container {
	return nbtContainer{NewCompoundTag()}
}

// Accessor

//type Accessor struct {
//}
//
//func (Accessor) GetByte(tag *ByteTag) int8 {
//	return tag.raw
//}
//
//func (Accessor) SetByte(tag *ByteTag, value int8) {
//	tag.raw = value
//}
