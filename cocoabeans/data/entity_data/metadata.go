package entity_data

import (
	"Learning2/cocoabeans/data/particle_data"
	"Learning2/cocoabeans/data/world_data"
	"Learning2/cocoabeans/protocol/packet/packet_data"
)

type TypeId packet_data.Enum[TypeId]

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

type Metadata interface {
	packet_data.Byte | packet_data.VarInt | packet_data.Float | packet_data.String | packet_data.Chat | OptChat | packet_data.Slot | packet_data.Boolean | world_data.Rotation | packet_data.Position | OptPosition | world_data.Direction | OptUUID | OptBlockID | packet_data.NBTTag | particle_data.Particle | VillagerData | OptVarInt | Pose
}

type OptChat struct {
	Present packet_data.Boolean
	Value   packet_data.Chat
}
type OptPosition struct {
	Present packet_data.Boolean
	Value   packet_data.Position
}
type OptUUID struct {
	Present packet_data.Boolean
	Value   packet_data.UUID
}
type OptVarInt packet_data.VarInt
type OptBlockID packet_data.VarInt
