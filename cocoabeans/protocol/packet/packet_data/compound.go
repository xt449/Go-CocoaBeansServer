package packet_data

import (
	"Learning2/cocoabeans/data/entity_data"
)

type EntityMetadata[M entity_data.Metadata] struct {
	Index UnsignedByte // If Index is 0xFF then it is the end of the EntityMetadata
	Type  VarInt
	Value M
}
type Slot struct {
	Present   Boolean
	ItemId    VarInt // Optional (ommited when Present is False)
	ItemCount Byte   // Optional (ommited when Present is False)
	NBT       NBTTag // Optional (ommited when Present is False)
}
