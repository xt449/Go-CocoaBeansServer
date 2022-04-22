package packet_data

type OptChat struct {
	Present Boolean
	Value   Chat
}
type OptPosition struct {
	Present Boolean
	Value   Position
}
type OptUUID struct {
	Present Boolean
	Value   UUID
}
type OptVarInt VarInt
type OptBlockID VarInt

type Metadata interface {
	Byte | VarInt | Float | String | Chat | OptChat | Slot | Boolean | Rotation | Position | OptPosition | Direction | OptUUID | OptBlockID | NBTTag | Particle | VillagerData | OptVarInt | Pose
}
