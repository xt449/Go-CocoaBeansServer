package packet_data

type TypeId Enum[TypeId]

const (
	BYTE         TypeId = iota
	VARINT       TypeId = iota
	FLOAT        TypeId = iota
	STRING       TypeId = iota
	CHAT         TypeId = iota
	OPTCHAT      TypeId = iota
	SLOT         TypeId = iota
	BOOLEAN      TypeId = iota
	ROTATION     TypeId = iota
	POSITION     TypeId = iota
	OPTPOSITION  TypeId = iota
	DIRECTION    TypeId = iota
	OPTUUID      TypeId = iota
	OPTBLOCKID   TypeId = iota
	NBT          TypeId = iota
	PARTICLE     TypeId = iota
	VILLAGERDATA TypeId = iota
	OPTVARINT    TypeId = iota
	POSE         TypeId = iota
)
