package packet_data

import "CocoaBeansServer/cocoabeans/nbt"

type NBTTag nbt.Container
type Position struct {
	X int32
	Y int16
	Z int32
}
type Array[T interface{}] []T
type Enum[T interface{}] uint
type ByteArray []int8
