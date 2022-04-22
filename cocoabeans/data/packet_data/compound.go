package packet_data

type EntityMetadata[M Metadata] struct {
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
